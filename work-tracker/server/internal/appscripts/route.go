package appscripts

import (
	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool) func(r *router.Router){
	appScriptService := NewAppScriptsService(db)
	appScriptHttpHandler := NewAppScriptsHandler(appScriptService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		authRouter.HandleFunc("/appscript", handler.MakeHTTPHandleFunc(appScriptHttpHandler.AppScriptHandler)).Methods("GET")
		authRouter.HandleFunc("/appscript", handler.MakeHTTPHandleFunc(appScriptHttpHandler.UpdateAppScriptHandler)).Methods("PATCH")
	}
}