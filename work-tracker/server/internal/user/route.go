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
	userService := NewUserService(db)

	// Handler(repo)
	userHttpHandler := NewUserHttpHandler(userService)

	return func(r *router.Router) {
		router := r.Mux
		authRouter := r.AuthUserSubrouter()

		router.HandleFunc("/signup", handler.MakeHTTPHandleFunc(userHttpHandler.SignUpHandler)).Methods("POST")
		router.HandleFunc("/signin", handler.MakeHTTPHandleFunc(userHttpHandler.SignInHandler)).Methods("POST")

		// starts with -> /auth/
		authRouter.HandleFunc("/signout", handler.MakeHTTPHandleFunc(userHttpHandler.SignOutHandler)).Methods("POST")

		// Authenticated User Subrouter
		// authUserSubrouter := r.AuthUserSubrouter()
	}
}