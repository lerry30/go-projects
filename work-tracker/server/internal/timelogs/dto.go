package timelogs

import (
	"time"
)

type TimeLogsRequest struct {
	Login time.Time `json:"login"`
	FirstBreakOut time.Time `json:"first-break-out"`
	FirstBreakIn time.Time `json:"first-break-in"`
	LunchOut time.Time `json:"lunch-out"`
	LunchIn time.Time `json:"lunch-in"`
	SecondBreakOut time.Time `json:"second-break-out"`
	SecondBreakIn time.Time `json:"second-break-in"`
	Logout time.Time `json:"logout"`
}

type TimeLogsResponse struct {
	ID int32 `json:"id"`
	Login string `json:"login"`
	FirstBreakOut string `json:"first-break-out"`
	FirstBreakIn string `json:"first-break-in"`
	LunchOut string `json:"lunch-out"`
	LunchIn string `json:"lunch-in"`
	SecondBreakOut string `json:"second-break-out"`
	SecondBreakIn string `json:"second-break-in"`
	Logout string `json:"logout"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}