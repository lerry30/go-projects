package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConnection struct {
	Pool *pgxpool.Pool
}

func (p *PostgresConnection) Close() {
	p.Pool.Close()
}

func NewPostgresConnection(strConn string) *PostgresConnection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(strConn)
	if err != nil {
		log.Fatalf("invalid database connection string: %v", err)
	}

	config.MaxConns = 20
	config.MinConns = 5

	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	config.AfterConnect = func(ctx context.Context, pgx *pgx.Conn) error {
		log.Println("new connection established")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("unable to create pool %v\n", err)
	}

	return &PostgresConnection{
		Pool: pool,
	}
}