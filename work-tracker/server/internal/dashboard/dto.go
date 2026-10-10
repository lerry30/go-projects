package dashboard

import (
	"time"
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
	FirstBreakIn time.Time `json:"time-first-break-in"`
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