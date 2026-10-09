package dashboard

import (
	"fmt"
	"net/http"

	"tracker/internal/shared/account"
	"tracker/internal/shared/httpresponse"
)

type DashboardHttpHandler struct {
	dashboardRepo DashboardRepository
}

func NewDashboardHttpHandler(dashboardRepo DashboardRepository) *DashboardHttpHandler {
	return &DashboardHttpHandler{
		dashboardRepo: dashboardRepo,
	}
}

func (d *DashboardHttpHandler) DashboardHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	var resError httpresponse.ErrorResponse

	userID, err := account.GetCurrentUserID(r)
	if err != nil {
		resError.WriteMessage(http.StatusBadRequest, "user not found")
		return resError
	}

	allData, err := d.dashboardRepo.GetAll(userID)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%v", allData)
//
	//var resOK httpresponse.OKResponse
	//resOK.WriteMessage(http.StatusOK, "dashboard")
//
	//return resOK

	var res DashboarResponse
	res.WriteResponse(
		http.StatusOK,
		allData.FirstName,
		allData.LastName,
		allData.Username,
		allData.UserCreatedAt,
		allData.UserUpdatedAt,
		allData.Login,
		allData.FirstBreakOut,
		allData.FirstBreakIN,
		allData.LunchOut,
		allData.LunchIn,
		allData.SecondBreakOut,
		allData.SecondBreakIn,
		allData.Logout,
		allData.TimeCreatedAt,
		allData.TimeUpdatedAt,
		allData.TrackerSpreadSheetID,
		allData.TrackerYear,
		allData.TrackerMonth,
		allData.TrackerCreatedAt,
		allData.TrackerUpdatedAt,
	)

	return res
}