package handler

import (
	"net/http"

	"tracker/internal/shared/httpresponse"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) httpresponse.Response

func MakeHTTPHandleFunc(f APIFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := f(w, r)
		res.SendResponse(w)
	}
}