package tracker

import (
	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"

	"tracker/internal/datachange"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool, notif *datachange.NotifyScheduler) func(r *router.Router) {
	trackerService := NewTrackerService(db, notif)
	trackerHttpHandler := NewTrackerHttpHandler(trackerService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		// starts with -> /auth/
		authRouter.HandleFunc("/tracker", handler.MakeHTTPHandleFunc(trackerHttpHandler.TrackerHandler)).Methods("GET")
		authRouter.HandleFunc("/tracker", handler.MakeHTTPHandleFunc(trackerHttpHandler.UpdateTrackerHandler)).Methods("PATCH")
	}
}