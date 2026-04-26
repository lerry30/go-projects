package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PostgresAddressRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewPostgresAddressRepository(db *pgxpool.Pool, rdb *redis.Client) *PostgresAddressRepository {
	return &PostgresAddressRepository{db: db, rdb: rdb}
}

func (r *PostgresAddressRepository) Create(addr map[string]any, userID int32) (*addressAggre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to add address: \n %w", err)
	}
	defer tx.Rollback(ctx) // no-op if already committed

	// check if country is available

	country, ok := addr["country_name"]
	if !ok {
		return nil, fmt.Errorf("Invalid country name\n")
	}
	delete(addr, "country_name")

	ctr, ok := country.(string)
	if !ok {
		return nil, fmt.Errorf("Invalid country name\n")
	}

	countryRes, err := QueryRow[string, countryRow](tx, "country", "country_name", ctr)
	if err != nil {
		return nil, fmt.Errorf("Country is not available\n")
	}

	// create new address

	addr["country_id"] = countryRes.ID

	addrRes, err := Create[addressRow](tx, "address", addr)
	if err != nil {
		return nil, fmt.Errorf("Failed to add new address: \n %w", err)
	}

	// associate address to user

	usrAddr := pgx.NamedArgs{
		"user_id":    userID,
		"address_id": addrRes.ID,
	}
	userAddrRes, err := Create[userAddressRow](tx, "user_address", usrAddr)
	if err != nil {
		return nil, fmt.Errorf("Failed to associate new address to user: \n %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("Failed to add adress: \n %w", err)
	}

	return &addressAggre{
		ID:          addrRes.ID,
		Line1:       addrRes.Line1,
		Line2:       addrRes.Line2,
		City:        addrRes.City,
		Region:      addrRes.Region,
		ZIPCode:     addrRes.ZIPCode,
		CountryName: countryRes.CountryName,
		IsDefault:   userAddrRes.IsDefault,
		CreatedAt:   addrRes.CreatedAt,
		UpdatedAt:   userAddrRes.UpdatedAt,
	}, nil
}

func (r *PostgresAddressRepository) GetAddresses(userID, limit, offset int32) (*[]addressAggreTtl, error) {
	arg := pgx.NamedArgs{"user_id": userID, "limit": limit, "offset": offset}
	query := `
		SELECT
			a.id,
			a.address_line1,
			a.address_line2,
			a.city,
			a.region,
			a.zip_code,
			c.country_name,
			ua.is_default,
			a.created_at,
			ua.updated_at,
			COUNT(*) OVER() AS total_count
		FROM address a
		JOIN country c ON a.country_id = c.id
		JOIN user_address ua ON a.id = ua.address_id
		WHERE ua.user_id = @user_id
		ORDER BY id DESC
		LIMIT @limit OFFSET @offset;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query, arg)
	if err != nil {
		return nil, fmt.Errorf("Execute database query create error: %w", err)
	}
	defer rows.Close()

	userAddresses, err := pgx.CollectRows(rows, pgx.RowToStructByName[addressAggreTtl])
	if err != nil {
		return nil, fmt.Errorf("CollectRows to addressAggre error: %w", err)
	}

	return &userAddresses, nil
}

func (r *PostgresAddressRepository) Delete(addrID int32) (int32, error) {
	query := `
		WITH deleted_address AS (
			DELETE FROM user_address
			WHERE address_id = $1
		)
		DELETE FROM address
		WHERE id = $1
		RETURNING id;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int32
	err := r.db.QueryRow(ctx, query, addrID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Failed to delete row in user_address and address: %w", err)
	}

	return id, nil
}

func (r *PostgresAddressRepository) SetDefault(addrID int32, userID int32) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("Failed to set address as default: %w", err)
	}
	defer tx.Rollback(ctx)

	queryOff := `
		UPDATE user_address 
		SET is_default=false 
		WHERE user_id=$1 AND is_default=true;
	`

	querySet := `
		UPDATE user_address 
		SET is_default=true 
		WHERE user_id=$1 AND address_id=$2
		RETURNING address_id;
	`

	// set current default address to false
	_, err = tx.Exec(ctx, queryOff, userID)
	if err != nil {
		return 0, fmt.Errorf("Failed to set default address. Address cannot be changed to regular: %w", err)
	}

	// set selected address as default
	var id int32
	err = tx.QueryRow(ctx, querySet, userID, addrID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Failed to set address as default: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("Failed to set address as default. Rolling back changes: %w", err)
	}

	return id, nil
}
