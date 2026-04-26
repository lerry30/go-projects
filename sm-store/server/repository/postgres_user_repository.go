package repository

import (
	"fmt"
	"time"

	"small-store/cache"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PostgresUserRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewPostgresUserRepository(db *pgxpool.Pool, rdb *redis.Client) *PostgresUserRepository {
	return &PostgresUserRepository{db: db, rdb: rdb}
}

func (r *PostgresUserRepository) GetByID(id int32) (*userRow, error) {
	cacheKey := fmt.Sprintf("user%d", id)

	// redis get
	cacheUsr, err := cache.Get[userRow](r.rdb, cacheKey)
	if err == nil {
		return &cacheUsr, nil // cache hit
	}

	// database get
	userRes, err := QueryRow[int32, userRow](r.db, "site_user", "id", id)
	if err != nil {
		return nil, fmt.Errorf("DB Query user error:\n %w", err)
	}

	// redis save
	go func(u userRow) {
		u.Password = ""
		if err := cache.Set(r.rdb, cacheKey, u, 20*time.Minute); err != nil {
			fmt.Printf("failed to cache user data: %s", err)
		}
	}(userRes) // pass by value avoids data race on the userRest variable after the parent function returns

	return &userRes, nil
}

func (r *PostgresUserRepository) Create(user map[string]any) (int32, error) {
	userRes, err := Create[userRow](r.db, "site_user", user)
	if err != nil {
		return 0, fmt.Errorf("Create Error: \n %w", err)
	}

	return userRes.ID, nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*userRow, error) {
	userRes, err := QueryRow[string, userRow](r.db, "site_user", "email_address", email)
	if err != nil {
		return nil, fmt.Errorf("DB Query user email error:\n %w", err)
	}

	return &userRes, nil
}

func (r *PostgresUserRepository) Update(user map[string]any, id int32) (*userRow, error) {
	userRes, err := Update[int32, userRow](r.db, "site_user", user, "id", id)
	if err != nil {
		return nil, fmt.Errorf("Failed to update user: \n %w", err)
	}

	// delete redis data
	cacheKey := fmt.Sprintf("user%d", userRes.ID)
	cache.Delete(r.rdb, cacheKey)

	return &userRes, nil
}
