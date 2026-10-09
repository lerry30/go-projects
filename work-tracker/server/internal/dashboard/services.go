package dashboard

import (
	"fmt"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

type DashboardService struct {
	db *pgxpool.Pool
}

func NewDashboardService(db *pgxpool.Pool) *DashboardService {
	return &DashboardService{
		db: db,
	}
}

func (d *DashboardService) GetAll(userID int64) (*EmployeeTimeLogsEntity, error) {
	args := pgx.NamedArgs{"user_id": userID}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := d.db.Query(ctx, queryAllStmt, args)
	if err != nil {
		return nil, fmt.Errorf("database query operation failed: %w", err)
	}
	defer rows.Close()

	dashboardData, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[EmployeeTimeLogsEntity])
	if err != nil {
		return nil, fmt.Errorf("database query result does not match the expected entity: %w", err)
	}

	return &dashboardData, nil
}