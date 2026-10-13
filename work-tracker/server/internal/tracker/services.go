package tracker

import (
	"fmt"

	"tracker/internal/platform/postgres/gen_queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TrackerService struct {
	db *pgxpool.Pool
}

func NewTrackerService(db *pgxpool.Pool) *TrackerService {
	return &TrackerService{
		db: db,
	}
}

func (t *TrackerService) Tracker(userID int32) (*TrackerEntity, error) {
	dbTracker, err := gen_queries.QueryRow[int32, TrackerEntity](t.db, TableName, TrackerCols.EmployeeID, userID)
	if err != nil {
		return nil, fmt.Errorf("db query tracker not found: %w", err)
	}
	return &dbTracker, nil
}