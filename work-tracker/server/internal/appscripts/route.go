package appscripts

import (
	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"

	"tracker/internal/datachange"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool, notif *datachange.NotifyScheduler) func(r *router.Router){
	appScriptService := NewAppScriptsService(db, notif)
	appScriptHttpHandler := NewAppScriptsHttpHandler(appScriptService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		authRouter.HandleFunc("/appscript", handler.MakeHTTPHandleFunc(appScriptHttpHandler.AppScriptHandler)).Methods("GET")
		authRouter.HandleFunc("/appscript-endpoint", handler.MakeHTTPHandleFunc(appScriptHttpHandler.UpdateAppScriptEndpointHandler)).Methods("PATCH")
		authRouter.HandleFunc("/appscript-status", handler.MakeHTTPHandleFunc(appScriptHttpHandler.UpdateAppScriptStatusHandler)).Methods("PATCH")
	}
}