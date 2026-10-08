package dashboard

import (
	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool) func(r *router.Router) {
	// repository
	dshbrdService := NewDashboardService(db)

	// handler
	dshbrdRepo := NewDashboardHttpHandler(dshbrdService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		// starts with -> /auth/
		authRouter.HandleFunc("/dashboard", handler.MakeHTTPHandleFunc(dshbrdRepo.DashboardHandler)).Methods("GET")
	}
}