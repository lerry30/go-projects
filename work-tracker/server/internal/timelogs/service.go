package timelogs

import (
	"fmt"

	"tracker/internal/platform/postgres/gen_queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TimeLogsService struct {
	db *pgxpool.Pool
}

func NewTimeLogsService(db *pgxpool.Pool) *TimeLogsService {
	return &TimeLogsService{
		db: db,
	}
}

func (t *TimeLogsService) TimeLogs(userID int32) (*TimeLogsEntity, error) {
	dbTimeLogs, err := gen_queries.QueryRow[int32, TimeLogsEntity](t.db, TableName, TimeLogsCols.EmployeeID, userID)
	if err != nil {
		return nil, fmt.Errorf("db query time_logs error: %w", err)
	}
	return &dbTimeLogs, nil
}