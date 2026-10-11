package appscripts

import (
	"time"
)

// --+ request
type AppScriptsRequest struct {
	ID int32 `json:"id"`
	Endpoint string `json:"endpoint"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}

// --+ response
type AppScriptsResponse struct {
	ID int32 `json:"id"`
	Endpoint string `json:"endpoint"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}