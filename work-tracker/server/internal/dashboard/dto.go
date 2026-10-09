package dashboard

import (
	"time"
	"encoding/json"
	"net/http"
)

// --+ request

// --+ response
type DashboardDataResponse struct {
	FirstName string `json:"first-name"`
	LastName string `json:"last-name"`
	Username string `json:"username"`
	UserCreatedAt time.Time `json:"user-created-at"`
	UserUpdatedAt time.Time `json:"user-updated-at"`
	Login time.Time `json:"time-login"`
	FirstBreakOut time.Time `json:"time-first-break-out"`
	FirstBreakIN time.Time `json:"time-first-break-in"`
	LunchOut time.Time `json:"time-lunch-out"`
	LunchIn time.Time `json:"time-lunch-in"`
	SecondBreakOut time.Time `json:"time-second-break-out"`
	SecondBreakIn time.Time `json:"time-second-break-in"`
	Logout time.Time `json:"time-logout"`
	TimeCreatedAt time.Time `json:"time-created-at"`
	TimeUpdatedAt time.Time `json:"time-updated-at"`
	TrackerSpreadSheetID string `json:"tracker-spread-sheet-id"`
	TrackerYear int16 `json:"tracker-year"`
	TrackerMonth int16 `json:"tracker-month"`
	TrackerCreatedAt time.Time `json:"tracker-created-at"`
	TrackerUpdatedAt time.Time `json:"tracker-updated-at"`
}

type DashboarResponse struct {
	DashboardDataResponse
	HttpCode int
}

func (d *DashboarResponse) WriteResponse(
		statusCode int,
		firstName string, 
		lastName string, 
		username string, 
		userCreatedAt time.Time, 
		userUpdatedAt time.Time, 
		login time.Time, 
		firstBreakOut time.Time, 
		firstBreakIN time.Time, 
		lunchOut time.Time, 
		lunchIn time.Time, 
		secondBreakOut time.Time, 
		secondBreakIn time.Time, 
		logout time.Time, 
		timeCreatedAt time.Time,
		timeUpdatedAt time.Time,
		trackerSpreadSheetID string, 
		trackerYear int16, 
		trackerMonth int16, 
		trackerCreatedAt time.Time, 
		trackerUpdatedAt time.Time, 
	) {
	d.HttpCode = statusCode
	d.DashboardDataResponse = DashboardDataResponse{
		FirstName: firstName,
		LastName: lastName,
		Username: username,
		UserCreatedAt: userCreatedAt,
		UserUpdatedAt: userUpdatedAt,
		Login: login,
		FirstBreakOut: firstBreakOut,
		FirstBreakIN: firstBreakIN,
		LunchOut: lunchOut,
		LunchIn: lunchIn,
		SecondBreakOut: secondBreakOut,
		SecondBreakIn: secondBreakIn,
		Logout: logout,
		TimeCreatedAt: timeCreatedAt,
		TimeUpdatedAt: timeUpdatedAt,
		TrackerSpreadSheetID: trackerSpreadSheetID,
		TrackerYear: trackerYear,
		TrackerMonth: trackerMonth,
		TrackerCreatedAt: trackerCreatedAt,
		TrackerUpdatedAt: trackerUpdatedAt,
	}
}

func (d DashboarResponse) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(d.HttpCode)
	json.NewEncoder(w).Encode(d.DashboardDataResponse)
}