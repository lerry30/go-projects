package repository

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Primitive interface {
	int | int32 | string | bool
}

func QueryRow[T Primitive, R any](dbPool *pgxpool.Pool, table string, column string, value T) (R, error) {
	var zero R

	if err := ValidateIdentifier(table); err != nil {
		return zero, fmt.Errorf("invalid table name: %w", err)
	}
	if err := ValidateIdentifier(column); err != nil {
		return zero, fmt.Errorf("invalid table column name: %w", err)
	}

	ph := fmt.Sprintf("@%s", column)
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=%s", table, column, ph)

	arg := make(pgx.NamedArgs, 1)
	arg[column] = value

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := dbPool.Query(ctx, query, arg)
	if err != nil {
		return zero, fmt.Errorf("failed to execute query by column: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[R]) // R, error
}

func Create[R any](dbPool *pgxpool.Pool, table string, args pgx.NamedArgs) (R, error) {
	var zero R
	columns := slices.Collect(maps.Keys(args))

	// Whitelist validation to prevent SQL injection
	if err := ValidateIdentifier(table); err != nil {
		return zero, fmt.Errorf("invalid table name: %w", err)
	}
	for _, col := range columns {
		if err := ValidateIdentifier(col); err != nil {
			return zero, fmt.Errorf("invalid column name %q: %w", col, err)
		}
	}

	// Build placeholders: (@name, @email, $password, ...)
	placeholders := make([]string, len(columns))
	for i, col := range columns {
		placeholders[i] = fmt.Sprintf("@%s", col)
	}

	// Build column list: (col1, col2, col3)
	cols := strings.Join(columns, ", ")
	ph := strings.Join(placeholders, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING *;", table, cols, ph)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := dbPool.Query(ctx, query, args)
	if err != nil {
		return zero, fmt.Errorf("failed to execute insert: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[R])
}

func Update[T Primitive, R any](dbPool *pgxpool.Pool, table string, args pgx.NamedArgs, column string, value T) (R, error) {
	var zero R
	columns := slices.Collect(maps.Keys(args))

	// Whitelist validation to prevent SQL injection
	if err := ValidateIdentifier(table); err != nil {
		return zero, fmt.Errorf("invalid table name: %w", err)
	}
	for _, col := range columns {
		if err := ValidateIdentifier(col); err != nil {
			return zero, fmt.Errorf("invalid column name %q: %w", col, err)
		}
	}

	// Build column = placeholders: (name = @name, email = @email, password = @password, ...)
	sqlParams := make([]string, len(columns))
	for i, col := range columns {
		sqlParams[i] = fmt.Sprintf("%s=@%s", col, col)
	}

	pr := strings.Join(sqlParams, ", ")
	phCol := fmt.Sprintf("@%s", column)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s=%s RETURNING *;", table, pr, column, phCol)

	// where clause value
	args[column] = value

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := dbPool.Query(ctx, query, args)
	if err != nil {
		return zero, fmt.Errorf("failed to execute insert: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[R])
}
