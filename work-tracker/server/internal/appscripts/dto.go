package appscripts

import (
	"time"
)

// --+ request
type AppScriptsEndpointRequest struct {
	Endpoint string `json:"endpoint"`
}

type AppScriptsStatusRequest struct {
	Status string `json:"status"`
}

// --+ response
type AppScriptsResponse struct {
	ID int32 `json:"id"`
	Endpoint string `json:"endpoint"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}