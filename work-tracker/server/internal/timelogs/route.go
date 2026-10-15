package timelogs

import (

	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool) func(r *router.Router) {
	timeLogsService := NewTimeLogsService(db)
	timeLogsHttpHandler := NewTimeLogsHttpHandler(timeLogsService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		authRouter.HandleFunc("/time-logs", handler.MakeHTTPHandleFunc(timeLogsHttpHandler.TimeLogsHandler)).Methods("GET")
	}
}