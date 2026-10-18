package appscripts

import (
	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool) func(r *router.Router){
	appScriptService := NewAppScriptsService(db)
	appScriptHttpHandler := NewAppScriptsHttpHandler(appScriptService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		authRouter.HandleFunc("/appscript", handler.MakeHTTPHandleFunc(appScriptHttpHandler.AppScriptHandler)).Methods("GET")
		authRouter.HandleFunc("/appscript-endpoint", handler.MakeHTTPHandleFunc(appScriptHttpHandler.UpdateAppScriptEndpointHandler)).Methods("PATCH")
		authRouter.HandleFunc("/appscript-status", handler.MakeHTTPHandleFunc(appScriptHttpHandler.UpdateAppScriptStatusHandler)).Methods("PATCH")
	}
}