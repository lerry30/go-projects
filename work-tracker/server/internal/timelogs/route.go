package timelogs

import (

	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"

	"tracker/internal/datachange"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool, notif *datachange.NotifyScheduler) func(r *router.Router) {
	timeLogsService := NewTimeLogsService(db, notif)
	timeLogsHttpHandler := NewTimeLogsHttpHandler(timeLogsService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		authRouter.HandleFunc("/time-logs", handler.MakeHTTPHandleFunc(timeLogsHttpHandler.TimeLogsHandler)).Methods("GET")
		authRouter.HandleFunc("/time-logs", handler.MakeHTTPHandleFunc(timeLogsHttpHandler.UpdateTimeLogsHandler)).Methods("PATCH")
	}
}