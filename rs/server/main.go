package main

import (
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// API server
	var port string = ":" + os.Getenv("PORT")
	router := mux.NewRouter()
}
