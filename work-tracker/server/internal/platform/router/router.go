package router

import (
	"tracker/internal/shared/middleware"
	"tracker/internal/shared/middleware/jwt"

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

	return &Router{
		Mux: router,
	}
}

func (r *Router) AuthUserSubrouter() *mux.Router {
	subrouter := r.Mux.PathPrefix("/auth/").Subrouter()
	subrouter.Use(jwt.JWT)
	return subrouter
}