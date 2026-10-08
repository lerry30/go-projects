package dashboard

import (
	"net/http"

	"tracker/internal/shared/httpresponse"
)

type DashboardHttpHandler struct {
	dshbrdRepo DashboardRepository
}

func NewDashboardHttpHandler(dshbrdRepo DashboardRepository) *DashboardHttpHandler {
	return &DashboardHttpHandler{
		dshbrdRepo: dshbrdRepo,
	}
}

func (d *DashboardHttpHandler) DashboardHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {

	var resOK httpresponse.OKResponse
	resOK.WriteMessage(http.StatusOK, "dashboard")

	return resOK
}