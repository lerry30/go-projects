package repository

import (
	"time"
)

// --- DB row types (1:1 table mapping) ---

type countryRow struct {
	ID          int32  `db:"id"`
	CountryName string `db:"country_name"`
}

type addressRow struct {
	ID        int32     `db:"id"`
	Line1     string    `db:"address_line1"`
	Line2     string    `db:"address_line2"`
	City      string    `db:"city"`
	Region    string    `db:"region"`
	ZIPCode   int16     `db:"zip_code"`
	CountryID int32     `db:"country_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type userAddressRow struct {
	UserID    int32     `db:"user_id"`
	AddressID int32     `db:"address_id"`
	IsDefault bool      `db:"is_default"`
	UpdatedAt time.Time `db:"updated_at"`
}

// --- Aggregate/domain types (multi-table joins) ---

type addressAggre struct {
	ID          int32     `db:"id"`
	Line1       string    `db:"address_line1"`
	Line2       string    `db:"address_line2"`
	City        string    `db:"city"`
	Region      string    `db:"region"`
	ZIPCode     int16     `db:"zip_code"`
	CountryName string    `db:"country_name"`
	IsDefault   bool      `db:"is_default"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type addressAggreTtl struct {
	addressAggre
	TotalCount int64 `db:"total_count"`
}
