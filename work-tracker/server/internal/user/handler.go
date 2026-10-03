package user

import (
	"strings"

	"net/http"
	"encoding/json"

	"tracker/internal/shared/httpresponse"
	"tracker/internal/shared/utils"
)

type UserHttpHandler struct {
	usrRepo UserRepository
}

func NewUserHttpHandler(usrRepo UserRepository) *UserHttpHandler {
	return &UserHttpHandler{
		usrRepo: usrRepo,
	}
}

func (u *UserHttpHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	defer r.Body.Close()

	var userReq userRequest

	var resError httpresponse.ErrorResponse

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		resError.WriteMessage(http.StatusBadRequest, "invalid JSON")
		return resError
	}

	userReq.FirstName = utils.Capitalize(userReq.FirstName)
	userReq.LastName = utils.Capitalize(userReq.LastName)
	userReq.Username = strings.TrimSpace(userReq.Username)
	userReq.Password = strings.TrimSpace(userReq.Password)

	if userReq.FirstName == "" ||
	userReq.LastName == "" ||
	userReq.Username == "" ||
	userReq.Password == "" {
		resError.WriteMessage(http.StatusBadRequest, "missing require fields")
		return resError
	}

	// TODO: filter user input values (e.g., validate characters)

	dbUser, err := u.usrRepo.GetByUsername(userReq.Username)
	if dbUser != nil || err == nil {
		resError.WriteMessage(http.StatusConflict, "record already exists")
		return resError
	}

	token, err := u.usrRepo.Create(userReq)
	if err != nil {
		resError.WriteMessage(http.StatusBadRequest, "failed to signup user")
		return resError
	}

	w.Header().Set("Authorization", "Bearer "+token)

	var resOK httpresponse.OKResponse
	resOK.WriteMessage(http.StatusOK, "User successfully signed up")
	return resOK
}