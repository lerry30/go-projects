package httpresponse

import (
	"net/http"
	"encoding/json"
)

type DataResponse struct {
	Body any
	HttpCode int
}

func (d *DataResponse) WriteBody(statusCode int, data any) {
	d.HttpCode = statusCode
	d.Body = data
}

func (d DataResponse) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(d.HttpCode)
	json.NewEncoder(w).Encode(d.Body)
}