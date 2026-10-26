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
	"tracker/internal/users"
	"tracker/internal/appscripts"
	"tracker/internal/tracker"
	"tracker/internal/timelogs"

	"tracker/internal/in_process_scheduling"

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

	//in_process_scheduling.RunSchedule(db.Pool)
	loc, _ := time.LoadLocation("Asia/Manila")
	sched, err := in_process_scheduling.NewScheduler(loc, db.Pool)
	if err != nil {
		log.Fatal(err)
	}
	sched.Start()

	// -----
	
	var rt *router.Router = router.NewRouter()

	// HTTP enpoint registration
	registrars := []RouteRegistrar{
		users.Route(db.Pool),
		appscripts.Route(db.Pool),
		tracker.Route(db.Pool),
		timelogs.Route(db.Pool),
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
	ctxQuit, quit := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer quit()
	<-ctxQuit.Done()
	// blocks here until CTRL+C or system termination

	// Shutdown sequence (fully reachable now)
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Gracefully shutdown the server (waits for active requests to finish)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	_ = sched.Stop(ctx)

	// Close Postgres and Redis connetion
	// rdb.Close()
	db.Close()

	log.Println("server exited cleanly")
}