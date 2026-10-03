package httpresponse

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct{Message string `json:"warning"`}

type ErrorResponse struct {
	ErrorMessage
	HttpCode int
}

func (e *ErrorResponse) WriteMessage(statusCode int, str string) {
	e.HttpCode = statusCode
	e.ErrorMessage = ErrorMessage{
		Message: str,
	}
}

func (e ErrorResponse) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HttpCode)
	json.NewEncoder(w).Encode(e.ErrorMessage)
}


// Usage:
/*
	res := ErrorResponse{}
	res.WriteMessage(http.StatusNotFound, "Not found")
*/