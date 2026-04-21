package repository

import (
	"time"
)

type UserRow struct {
	ID        int32     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email_address"`
	Phone     string    `db:"phone_number"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserRepository interface {
	GetByID(id string) (*UserRow, error)
	GetByEmail(email string) (*UserRow, error)
	Create(user map[string]any) (int32, error)
	Update(user map[string]any, id int32) (*UserRow, error)
}
