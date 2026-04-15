package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db *DB = NewDB(string(os.Getenv("DB_CONN_URL")))
	defer db.pool.Close()

	server := NewAPIServer(":8080", db)
	server.Run()
}
