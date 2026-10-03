package httpresponse

import (
	"net/http"
	"encoding/json"
)

type OKMessage struct{Message string `json:"message"`}

type OKResponse struct {
	OKMessage
	HttpCode int
}

func (k *OKResponse) WriteMessage(statusCode int, mssg string) {
	k.HttpCode = statusCode
	k.OKMessage = OKMessage{
		Message: mssg,
	}
}

func (k OKResponse) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(k.HttpCode)
	json.NewEncoder(w).Encode(k.OKMessage)
}