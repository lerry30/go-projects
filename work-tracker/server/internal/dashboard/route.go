package dashboard

import (
	"tracker/internal/platform/router"
	"tracker/internal/shared/handler"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Route(db *pgxpool.Pool) func(r *router.Router) {
	// repository
	dashboardService := NewDashboardService(db)

	// handler
	dashboarRepo := NewDashboardHttpHandler(dashboardService)

	return func(r *router.Router) {
		authRouter := r.AuthUserSubrouter()

		// starts with -> /auth/
		authRouter.HandleFunc("/dashboard", handler.MakeHTTPHandleFunc(dashboarRepo.DashboardHandler)).Methods("GET")
	}
}