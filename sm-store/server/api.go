package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"small-store/auth"

	"github.com/gorilla/mux"
)

type APIServer struct {
	port string
	db   *DB
}

// function type
type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string
}

// ----
// Initialize
func NewAPIServer(port string, db *DB) *APIServer {
	return &APIServer{
		port: port,
		db:   db,
	}
}

// Run the server
func (apiServer *APIServer) Run() {
	router := mux.NewRouter()

	// Global OPTIONS handler for all routes
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		EnableCORS(&w)
		w.WriteHeader(http.StatusNoContent)
	})

	//router.HandleFunc("/test", makeHTTPHandleFunc(apiServer.TestHandler)).Methods("GET")
	router.HandleFunc("/signup", makeHTTPHandleFunc(apiServer.SignUpHandler)).Methods("POST")
	router.HandleFunc("/signin", makeHTTPHandleFunc(apiServer.SignInHandler)).Methods("POST")
	router.HandleFunc("/logout", makeHTTPHandleFunc(apiServer.LogoutHandler)).Methods("POST")

	private := router.PathPrefix("/priv").Subrouter()
	private.Use(auth.Middleware)
	//private.HandleFunc("/test", makeHTTPHandleFunc(apiServer.TestHandler)).Methods("GET")

	fmt.Println("Server starting on port:", apiServer.port[1:])
	log.Fatal(http.ListenAndServe(apiServer.port, router))
}

// Wrapper to simplify the error response handling
func makeHTTPHandleFunc(f apiFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		EnableCORS(&w)

		if err := f(w, r); err != nil {
			ResponseJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

// JSON response function
func ResponseJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

// Allow CORS
func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true") // Required for credentials
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
}
