package appscripts

import (
	"fmt"

	"tracker/internal/platform/postgres/gen_queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppScriptsService struct {
	db *pgxpool.Pool
}

func NewAppScriptsService(db *pgxpool.Pool) *AppScriptsService {
	return &AppScriptsService{
		db: db,
	}
}

func (a *AppScriptsService) AppScript(userID int32) (*AppScriptsEntity, error) {
	dbAppScript, err := gen_queries.QueryRow[int32, AppScriptsEntity](a.db, TableName, AppscriptsCols.EmployeeID, userID)
	if err != nil {
		return nil, fmt.Errorf("db query appscript not found: %w", err)
	}
	return &dbAppScript, nil
}