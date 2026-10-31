package tracker

import (
	"fmt"
	"context"
	"time"

	"tracker/internal/platform/postgres/gen_queries"

	"tracker/internal/datachange"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

type TrackerService struct {
	db *pgxpool.Pool
	notif *datachange.NotifyScheduler
}

func NewTrackerService(db *pgxpool.Pool, notif *datachange.NotifyScheduler) *TrackerService {
	return &TrackerService{
		db: db,
		notif: notif,
	}
}

func (t *TrackerService) Tracker(userID int32) (*TrackerEntity, error) {
	dbTracker, err := gen_queries.QueryRow[int32, TrackerEntity](t.db, TableName, TrackerCols.EmployeeID, userID)
	if err != nil {
		return nil, fmt.Errorf("db query tracker not found: %w", err)
	}
	return &dbTracker, nil
}

func (t *TrackerService) UpdateTracker(tracker *TrackerRequest, userID int32) (int32, error) {
	args := make(pgx.NamedArgs, 4)
	args[TrackerCols.EmployeeID] = userID
	args[TrackerCols.SpreadSheetID] = tracker.SpreadSheetID
	args[TrackerCols.Year] = tracker.Year
	args[TrackerCols.Month] = tracker.Month

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := t.db.Query(ctx, upsertTracker, args)
	if err != nil {
		return 0, fmt.Errorf("db upsert failed for tracker: %w", err)
	}
	defer rows.Close()

	dbTracker, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[TrackerEntity])
	if err != nil {
		return 0, fmt.Errorf("db failed to convert row to struct: %w", err)
	}

	if !t.notif.Trigger(userID) {
		return 0, fmt.Errorf("trigger failed: updating tracker")
	}
	return dbTracker.ID, nil
}