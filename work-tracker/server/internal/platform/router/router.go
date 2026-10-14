package router

import (
	"net/http"

	"tracker/internal/shared/middleware"

	"github.com/gorilla/mux"
)

/*
	I created a wrapper for the
	gorilla mux router to be able
	to create sub-router based on
	defined url prefix
*/

type Router struct {
	Mux *mux.Router
}

func NewRouter() *Router {
	router := mux.NewRouter()

	router.Use(middleware.CORS)
	router.Use(func(next http.Handler) http.Handler {
		// Middleware for cap limit that prevents large 
		// size of payload from the clients request. 1MB 
		// maximum number of bytes allowed.
		return middleware.LimitBody(next, 1<<20) // 1MB
	})

	return &Router{
		Mux: router,
	}
}

func (r *Router) AuthUserSubrouter() *mux.Router {
	subrouter := r.Mux.PathPrefix("/auth/").Subrouter()
	subrouter.Use(middleware.JWT)
	return subrouter
}