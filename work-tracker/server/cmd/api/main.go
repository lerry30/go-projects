package main

import (
	"log"
	"os"
	"context"
	"syscall"
	"time"
	"os/signal"
	"net/http"

	"tracker/internal/platform/postgres"
	/// "tracker/internal/platform/redis"
	"tracker/internal/platform/router"
	"tracker/internal/user"

	"github.com/joho/godotenv"
)

type RouteRegistrar func(r *router.Router)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("unable to access env variables")
	}

	// Environment variable declarations

	dbConnStr := os.Getenv("DATABASE_URL")

	// redisAddr := os.Getenv("UPSTASH_REDIS_ADDR")
	// redisPassword := os.Getenv("UPSTASH_REDIS_PASSWORD")

	var port string = ":" + os.Getenv("PORT")

	// -----

	var db *postgres.PostgresConnection = postgres.NewPostgresConnection(dbConnStr)
	defer db.Close()

	// var rdb *redis.RedisConnection = redis.NewRedisConnection(redisAddr, redisPassword)
	// defer rdb.Close()
	
	var rt *router.Router = router.NewRouter()

	// HTTP enpoint registration
	registrars := []RouteRegistrar{
		user.Route(db.Pool),
	}

	for _, registrar := range(registrars) {
		registrar(rt)
	}
	
	srv := &http.Server{
		Addr: port,
		Handler: rt.Mux,
	}

	go func() {
		log.Printf("server started on %s", port[1:])
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// ErrServerClosed is expected during graceful shutdown, not a real error
			log.Fatalf("server error: %v", err)
		}
	}()

	// Main goroutine is now free - sit here and wait for OS signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // blocks here until CTRL+C or system termination

	// Shutdown sequence (fully reachable now)
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shutdown the server (waits for active requests to finish)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	// Close Postgres and Redis connetion
	// rdb.Close()
	db.Close()

	log.Println("server exited cleanly")
}