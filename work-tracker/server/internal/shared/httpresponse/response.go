package httpresponse

import "net/http"

type Response interface {
	SendResponse(w http.ResponseWriter)
}