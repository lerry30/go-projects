package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"small-store/api"
	"small-store/auth"
	"small-store/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
	var db *DB = NewDB(string(os.Getenv("DB_CONN_URL")))
	defer db.pool.Close()

	// Redis
	rdb := NewRedisClient()
	defer rdb.Close() // clean up on exit

	// pattern: Interface-Based Dependency Injection
	userRepo := repository.NewPostgresUserRepository(db.pool, rdb)

	userHandler := api.NewUserHandler(userRepo)

	// API server
	var port string = ":" + os.Getenv("PORT")
	router := mux.NewRouter()

	// Global OPTIONS handler for all routes
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		EnableCORS(&w)
		w.WriteHeader(http.StatusNoContent)
	})

	router.HandleFunc("/signup", makeHTTPHandleFunc(userHandler.SignUpHandler)).Methods("POST")
	router.HandleFunc("/signin", makeHTTPHandleFunc(userHandler.SignInHandler)).Methods("POST")
	router.HandleFunc("/logout", makeHTTPHandleFunc(userHandler.LogoutHandler)).Methods("POST")

	private := router.PathPrefix("/priv").Subrouter()
	private.Use(auth.Middleware)

	private.HandleFunc("/user", makeHTTPHandleFunc(userHandler.ProfileHandler)).Methods("GET")
	private.HandleFunc("/me", makeHTTPHandleFunc(userHandler.UpdateUserHandler)).Methods("PATCH")

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		log.Printf("Server started on %s", port[1:])
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// ErrServerClosed is expected during graceful shutdown, not a real error
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Main goroutine is now FREE — sit here and wait for OS signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // blocks here until Ctrl+C or system termination

	// ---- Shutdown sequence (fully reachable now) ----
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully stop the HTTP server (waits for active requests to finish)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	// Close Redis, DB connections, etc.
	rdb.Close()
	db.pool.Close()

	log.Println("Server exited cleanly")
}
