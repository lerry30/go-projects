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

	key := string(os.Getenv("OPEN_WEATHER_API_KEY"))
	ow := NewOpenWeather(key)

	server := NewAPIServer(":8080")
	server.AddExternalAPI("open-weather", ow)
	server.Run()
}
