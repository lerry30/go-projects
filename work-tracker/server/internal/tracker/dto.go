package tracker

import (
	"time"
)

type TrackerResponse struct {
	ID int32 `json:"id"`
	SpreadSheetID string `json:"spread-sheet-id"`
	Year int16 `json:"year"`
	Month int16 `json:"month"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}