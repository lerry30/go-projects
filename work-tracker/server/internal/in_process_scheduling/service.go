package in_process_scheduling

import (
	"fmt"
	"context"
	"errors"
	"time"

	"tracker/internal/appscripts"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

type SchedulesService struct {
	db *pgxpool.Pool
}

func NewSchedulesService(db *pgxpool.Pool) *SchedulesService {
	return &SchedulesService{
		db: db,
	}
}

func (s *SchedulesService) RunningAppScripts() (*[]UsersScheduleEntity, error) {
	args := make(pgx.NamedArgs, 1)
	args[appscripts.AppScriptsCols.Status] = "running"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, selectUsersDetailsByStatus, args)
	if err != nil {
		return nil, fmt.Errorf("db error getting running appscripts: %w", err)
	}
	defer rows.Close()

	dbUsersSched, err := pgx.CollectRows(rows, pgx.RowToStructByName[UsersScheduleEntity])
	if errors.Is(err, pgx.ErrNoRows) {
		return &[]UsersScheduleEntity{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("faile to convert row to struct: %w", err)
	}
	
	return &dbUsersSched, nil
}