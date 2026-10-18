package appscripts

import (
	"fmt"
	"context"
	"time"

	"tracker/internal/platform/postgres/gen_queries"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
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
	dbAppScript, err := gen_queries.QueryRow[int32, AppScriptsEntity](a.db, TableName, AppScriptsCols.EmployeeID, userID)
	if err != nil {
		return nil, fmt.Errorf("db query appscript not found: %w", err)
	}
	return &dbAppScript, nil
}

func (a *AppScriptsService) UpdateAppScriptEndpoint(endpoint string, userID int32) (int32, error) {
	args := make(pgx.NamedArgs, 2)
	args[AppScriptsCols.EmployeeID] = userID
	args[AppScriptsCols.Endpoint] = endpoint

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := a.db.Query(ctx, upsertAppScriptURL, args)
	if err != nil {
		return 0, fmt.Errorf("db upsert failed for appscript: %w", err)
	}
	defer rows.Close()

	appScriptData, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[AppScriptsEntity])
	if err != nil {
		return 0, fmt.Errorf("db failed to convert row to struct: %w", err)
	}

	return appScriptData.ID, nil
}

func (a *AppScriptsService) UpdateAppScriptStatus(status string, userID int32) (int32, error) {
	args := make(pgx.NamedArgs, 1)
	args[AppScriptsCols.Status] = status

	dbAppScript, err := gen_queries.Update[int32, AppScriptsEntity](a.db, TableName, args, AppScriptsCols.EmployeeID, userID)
	if err != nil {
		return 0, fmt.Errorf("db error: %w", err)
	}
	return dbAppScript.ID, nil
}