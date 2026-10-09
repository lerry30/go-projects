package user

import "time"

type UsersEntity struct {
	ID int32 `db:"id"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	Username string `db:"username"`
	Password string `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

/*
var AllowedSortColumns = map[string]string{
	"first_name": "first_name",
	"last_name": "last_name",
	"created_at": "created_at",
	"updated_at": "updated_at",
}
*/

// Table column names
type UsersColumns struct {
	ID string
	FirstName string
	LastName string
	Username string
	Password string
	CreatedAt string
	UpdatedAt string
}

const TableName = "users"

var UserCols = UsersColumns{
	ID: "id",
	FirstName: "first_name",
	LastName: "last_name",
	Username: "username",
	Password: "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}