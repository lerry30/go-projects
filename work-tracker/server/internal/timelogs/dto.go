package timelogs

import (
	"time"

	"tracker/internal/shared/handler/types"
)

type TimeLogsRequest struct {
	Login types.ClockTime `json:"login"`
	FirstBreakOut types.ClockTime `json:"first-break-out"`
	FirstBreakIn types.ClockTime `json:"first-break-in"`
	LunchOut types.ClockTime `json:"lunch-out"`
	LunchIn types.ClockTime `json:"lunch-in"`
	SecondBreakOut types.ClockTime `json:"second-break-out"`
	SecondBreakIn types.ClockTime `json:"second-break-in"`
	Logout types.ClockTime `json:"logout"`
}

type TimeLogsResponse struct {
	ID int32 `json:"id"`
	Login types.ClockTime `json:"login"`
	FirstBreakOut types.ClockTime `json:"first-break-out"`
	FirstBreakIn types.ClockTime `json:"first-break-in"`
	LunchOut types.ClockTime `json:"lunch-out"`
	LunchIn types.ClockTime `json:"lunch-in"`
	SecondBreakOut types.ClockTime `json:"second-break-out"`
	SecondBreakIn types.ClockTime `json:"second-break-in"`
	Logout types.ClockTime `json:"logout"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}