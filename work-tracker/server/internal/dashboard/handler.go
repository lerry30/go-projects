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

	var res httpresponse.DataResponse

	resData := DashboardDataResponse{
		FirstName: allData.FirstName,
		LastName: allData.LastName,
		Username: allData.Username,
		UserCreatedAt: allData.UserCreatedAt,
		UserUpdatedAt: allData.UserUpdatedAt,
		Login: allData.Login,
		FirstBreakOut: allData.FirstBreakOut,
		FirstBreakIn: allData.FirstBreakIn,
		LunchOut: allData.LunchOut,
		LunchIn: allData.LunchIn,
		SecondBreakOut: allData.SecondBreakOut,
		SecondBreakIn: allData.SecondBreakIn,
		Logout: allData.Logout,
		TimeCreatedAt: allData.TimeCreatedAt,
		TimeUpdatedAt: allData.TimeUpdatedAt,
		TrackerSpreadSheetID: allData.TrackerSpreadSheetID,
		TrackerYear: allData.TrackerYear,
		TrackerMonth: allData.TrackerMonth,
		TrackerCreatedAt: allData.TrackerCreatedAt,
		TrackerUpdatedAt: allData.TrackerUpdatedAt,
	}

	res.WriteBody(http.StatusOK, resData)

	return res
}