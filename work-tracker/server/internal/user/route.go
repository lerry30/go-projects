package user

import (
	"tracker/internal/shared/handler"
	"tracker/internal/platform/router"

	"github.com/jackc/pgx/v5/pgxpool"
)

// func Route(db *pgxpool.Pool, rdb *redis.RedisConnection) func(r *router.Router){}

func Route(db *pgxpool.Pool) func(r *router.Router) {
	// Interface-based Repository Injection
	// Repo
	usrService := NewUserService(db)

	// Handler(repo)
	usrHttpHandler := NewUserHttpHandler(usrService)

	return func(r *router.Router) {
		router := r.Mux
		router.HandleFunc("/test", handler.MakeHTTPHandleFunc(usrHttpHandler.SignUpHandler)).Methods("POST")

		// Authenticated User Subrouter
		// authUserSubrouter := r.AuthUserSubrouter()
	}
}