package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"context"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "github.com/lib/pq"
)

func main() {
	// goose expects: migrate <command> [args]
	// e.g. go run ./cmd/migrate up
	//      go run ./cmd/migrate down
	//      go run ./cmd/migrate status
	//      go run ./cmd/migrate create add_users_table sql

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("error: failed to load env file")
	}

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("usage: migrate <command> [args] (e.g. up, down, status, create)")
	}

	command := args[0]

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}

	migrationDir := "migrations"

	switch command {
	case "create":
		if len(args) < 2 {
			log.Fatal("usage: migrate create <name> [sql|go]")
		}
		name := args[1]
		migType := "sql"
		if len(args) > 2 {
			migType = args[2]
		}
		if err := goose.Create(db, migrationDir, name, migType); err != nil {
			log.Fatalf("create failed: %v", err)
		}
	default:
		// up, down, status, redo, reset, version all route through goose.Run
		if err := goose.RunContext(context.Background(), command, db, migrationDir, args[1:]...); err != nil {
			log.Fatalf("migrate %s failed %v", command, err)
		}
	}

	fmt.Println("done")
}