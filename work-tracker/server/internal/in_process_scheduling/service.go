package in_process_scheduling

import (
	"fmt"
	"context"
	"errors"
	"time"

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

func (s *SchedulesService) Users() (*[]UsersScheduleEntity, error) {
	args := make(pgx.NamedArgs, 1)
	args[AppScriptsCols.Status] = "running"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, selectUsersDetailsByStatus, args)
	if err != nil {
		return nil, fmt.Errorf("db query failed in in_process_scheduling.Users(): %w", err)
	}
	defer rows.Close()

	dbUsersSched, err := pgx.CollectRows(rows, pgx.RowToStructByName[UsersScheduleEntity])
	if errors.Is(err, pgx.ErrNoRows) {
		return &[]UsersScheduleEntity{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to convert row to struct in in_process_scheduling.Users(): %w", err)
	}
	
	return &dbUsersSched, nil
}

func (s *SchedulesService) User(userID int32) (*UsersScheduleEntity, error) {
	args := make(pgx.NamedArgs, 1)
	args[AppScriptsCols.Status] = "running"
	args[UserCols.ID] = userID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, selectUser, args)
	if err != nil {
		return nil, fmt.Errorf("db query failed in in_process_scheduling.User(id): %w", err)
	}
	defer rows.Close()

	dbUserSched, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[UsersScheduleEntity])
	if errors.Is(err, pgx.ErrNoRows) {
		return &UsersScheduleEntity{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to convert row to struct in in_process_scheduling.User(id): %w", err)
	}
	
	return &dbUserSched, nil
}