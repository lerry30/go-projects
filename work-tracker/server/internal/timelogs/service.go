package timelogs

import (
	"fmt"
	"context"
	"time"

	"tracker/internal/platform/postgres/gen_queries"
	"tracker/internal/shared/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
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

func (t *TimeLogsService) UpdateTimeLogs(timeLogs *TimeLogsRequest, userID int32) (int32, error) {
	args := make(pgx.NamedArgs, 9)
	args[TimeLogsCols.EmployeeID] = userID
	args[TimeLogsCols.Login] = utils.To24HourFormat(timeLogs.Login.Time)
	args[TimeLogsCols.FirstBreakOut] = utils.To24HourFormat(timeLogs.FirstBreakOut.Time)
	args[TimeLogsCols.FirstBreakIn] = utils.To24HourFormat(timeLogs.FirstBreakIn.Time)
	args[TimeLogsCols.LunchOut] = utils.To24HourFormat(timeLogs.LunchOut.Time)
	args[TimeLogsCols.LunchIn] = utils.To24HourFormat(timeLogs.LunchIn.Time)
	args[TimeLogsCols.SecondBreakOut] = utils.To24HourFormat(timeLogs.SecondBreakOut.Time)
	args[TimeLogsCols.SecondBreakIn] = utils.To24HourFormat(timeLogs.SecondBreakIn.Time)
	args[TimeLogsCols.Logout] = utils.To24HourFormat(timeLogs.Logout.Time)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := t.db.Query(ctx, upsertTimeLogs, args)
	if err != nil {
		return 0, fmt.Errorf("failed to query timelogs: %w", err)
	}
	defer rows.Close()

	dbTimeLogs, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[TimeLogsEntity])
	if err != nil {
		return 0, fmt.Errorf("db failed to convert row to struct: %w", err)
	}

	return dbTimeLogs.ID, nil
}