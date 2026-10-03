package user

import (
	"fmt"

	"tracker/internal/platform/postgres/gen_queries"
	"tracker/internal/shared/utils"
	"tracker/internal/shared/middleware/jwt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) GetByUsername(username string) (*UsersEntity, error) {
	dbUser, err := gen_queries.QueryRow[string, UsersEntity](u.db, TableName, UserCols.Username, username)
	if err != nil {
		return nil, fmt.Errorf("db query user not found: %w", err)
	}
	return &dbUser, nil
}

func (u *UserService) Create(user userRequest) (string, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Printf("failed: %w", err)
		return "", fmt.Errorf("failed to encrypt user credentials: %w", err)
	}
	
	newRecord := make(pgx.NamedArgs, 4)
	newRecord[UserCols.FirstName] = user.FirstName
	newRecord[UserCols.LastName] = user.LastName
	newRecord[UserCols.Username] = user.Username
	newRecord[UserCols.Password] = hashedPassword
	
	newUser, err := gen_queries.Create[UsersEntity](u.db, TableName, newRecord)
	if err != nil {
		fmt.Printf("failed: %w", err)
		return "", fmt.Errorf("failed to create new user record: %w", err)
	}
	
	strID := fmt.Sprintf("%s", newUser.ID)
	token, err := jwt.GenerateToken(strID, newUser.Username)
	if err != nil {
		fmt.Printf("failed: %w", err)
		return "", fmt.Errorf("failed to create user token: %w", err)
	}

	return token, nil
}