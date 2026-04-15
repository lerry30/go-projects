package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(connStr string) *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Failed to parse config: %v\n", err)
	}

	// --- Pool sizing ---
	config.MaxConns = 20 // max open connections (default: 4 * NumCPU)
	config.MinConns = 5  // keep at least N connections warm

	// --- Timeouts & lifetimes ---
	config.MaxConnLifetime = 1 * time.Hour     // recycle connections after this duration
	config.MaxConnIdleTime = 30 * time.Minute  // close idle connections after this duration
	config.HealthCheckPeriod = 1 * time.Minute // how often to ping idle connections

	// --- Hook: run code after each new connection is established ---
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// e.g. register custom types, set session config
		log.Println("New connection established")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create pool: %v\n", err)
	}

	return &DB{
		pool: pool,
	}
}

// columns []string, args ...any — columns and values are now separate, typed, and matched by index. No more pre-built strings.
// validateIdentifier — since table/column names can't use $N params, this regex whitelists only safe identifiers before interpolation.
// Placeholders built dynamically — $1, $2, ... are generated to match the number of args, so pgx handles all value escaping.
// fmt.Errorf(...%w...) — original errors are wrapped and preserved for debugging.
//
// Example usage:
// goerr := db.Create(
//     "users",
//     []string{"name", "email", "age"},
//     "Alice", "alice@example.com", 30,
// )

// validateIdentifier rejects anything that isn't a simple snake_case/alphanumeric name.
// This prevents SQL injection via table/column names since those can't be parameterized.
var validIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func validateIdentifier(name string) error {
	if !validIdentifier.MatchString(name) {
		return fmt.Errorf("%q is not a valid SQL identifier", name)
	}
	return nil
}

func (db *DB) Create(table string, columns []string, args ...any) (string, error) {
	// Whitelist validation to prevent SQL injection
	if err := validateIdentifier(table); err != nil {
		return "", fmt.Errorf("invalid table name: %w", err)
	}
	for _, col := range columns {
		if err := validateIdentifier(col); err != nil {
			return "", fmt.Errorf("invalid column name %q: %w", col, err)
		}
	}

	// Build placeholders: ($1, $2, $3, ...)
	placeholders := make([]string, len(args))
	for i := range args {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	// Build column list: (col1, col2, col3)
	cols := strings.Join(columns, ", ")
	ph := strings.Join(placeholders, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id;", table, cols, ph)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id string
	err := db.pool.QueryRow(ctx, query, args...).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to execute insert: %w", err)
	}

	return id, nil
}

func (db *DB) Update(table string, columns []string, args ...any) (string, error) {
	if err := validateIdentifier(table); err != nil {
		return "", fmt.Errorf("invalid table name: %w", err)
	}

	for _, col := range columns {
		if err := validateIdentifier(col); err != nil {
			return "", fmt.Errorf("invalid column name %q: %w", col, err)
		}
	}

	// build column = placeholder
	placeholders := make([]string, len(columns))
	for i, col := range columns {
		placeholders[i] = fmt.Sprintf("%s=$%d", col, i+1)
	}

	colsPh := strings.Join(placeholders, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s RETURNING id;", table, colsPh)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id string
	err := db.pool.QueryRow(ctx, query, args...).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to execute update: %w", err)
	}

	return id, nil
}

func (db *DB) Delete(table, column, value string) (string, error) {
	if err := validateIdentifier(table); err != nil {
		return "", fmt.Errorf("invalid table name: %w", err)
	}

	if err := validateIdentifier(column); err != nil {
		return "", fmt.Errorf("invalid column name: %w", err)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1 RETURNING id;", table, column)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id string
	err := db.pool.QueryRow(ctx, query, value).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to execute delete: %w", err)
	}

	return id, nil
}

// Query instead of QueryRow to get all rows
// rows.FieldDescriptions() to dynamically read column count and names
// A pointer-indirection trick (valuePtrs) so Scan can populate values without knowing types at compile time
// rows.Err() check after the loop, which catches errors that interrupted iteration
func (db *DB) QueryAll(table string) ([]map[string]any, error) {
	if err := validateIdentifier(table); err != nil {
		return nil, fmt.Errorf("invalid table name: %w", err)
	}

	query := fmt.Sprintf("SELECT * FROM %s;", table)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query table %q: %w", table, err)
	}
	defer rows.Close()

	var results []map[string]any

	for rows.Next() {
		// Build a slice of destinations based on the number of columns
		fields := rows.FieldDescriptions()
		values := make([]any, len(fields))
		valuePtrs := make([]any, len(fields))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		row := make(map[string]any, len(fields))
		for i, field := range fields {
			row[string(field.Name)] = values[i]
		}
		results = append(results, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return results, nil
}

func (db *DB) QueryRow(table, column, value string) (map[string]any, error) {
	if err := validateIdentifier(table); err != nil {
		return nil, fmt.Errorf("invalid table name: %w", err)
	}

	if err := validateIdentifier(column); err != nil {
		return nil, fmt.Errorf("invalid column %q: %w", column, err)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1;", table, column)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.pool.Query(ctx, query, value)
	if err != nil {
		return nil, fmt.Errorf("failed to query table %q: %w", table, err)
	}
	defer rows.Close()

	result := make(map[string]any)

	for rows.Next() {
		fields := rows.FieldDescriptions()
		values := make([]any, len(fields))
		valuePtrs := make([]any, len(fields))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		for i, field := range fields {
			result[string(field.Name)] = values[i]
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return result, nil
}
