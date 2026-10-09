package user

import (
	"strings"

	"net/http"
	"encoding/json"

	"tracker/internal/shared/httpresponse"
	"tracker/internal/shared/utils"
	"tracker/internal/shared/token/jwt"
	"tracker/internal/shared/account"
)

type UserHttpHandler struct {
	userRepo UserRepository
}

func NewUserHttpHandler(userRepo UserRepository) *UserHttpHandler {
	return &UserHttpHandler{
		userRepo: userRepo,
	}
}

func (u *UserHttpHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	defer r.Body.Close()

	var userReq userSignUpRequest

	var resError httpresponse.ErrorResponse

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		resError.WriteMessage(http.StatusBadRequest, "invalid JSON")
		return resError
	}

	userReq.FirstName = utils.Capitalize(userReq.FirstName)
	userReq.LastName = utils.Capitalize(userReq.LastName)
	userReq.Username = strings.TrimSpace(userReq.Username)
	userReq.Password = strings.TrimSpace(userReq.Password)

	// DONE: update the empty validation to length based characters

	if len(userReq.FirstName) < 2 ||
	len(userReq.LastName) < 2 ||
	len(userReq.Username) < 4 ||
	len(userReq.Password) < 4 {
		resError.WriteMessage(http.StatusBadRequest, "Insufficient number of characters")
		return resError
	}

	// TODO: filter user input values (e.g., validate characters)

	dbUser, err := u.userRepo.GetByUsername(userReq.Username)
	if dbUser != nil || err == nil {
		resError.WriteMessage(http.StatusConflict, "record already exists")
		return resError
	}

	token, err := u.userRepo.Create(userReq)
	if err != nil {
		resError.WriteMessage(http.StatusBadRequest, "failed to signup user")
		return resError
	}

	w.Header().Set("Authorization", "Bearer "+token)

	var resOK httpresponse.OKResponse
	resOK.WriteMessage(http.StatusOK, "User successfully signed up")
	return resOK
}

func (u *UserHttpHandler) SignInHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	defer r.Body.Close()

	var userReq userSignInRequest

	var resError httpresponse.ErrorResponse

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		resError.WriteMessage(http.StatusBadRequest, "invalid JSON")
		return resError
	}

	userReq.Username = strings.TrimSpace(userReq.Username)
	userReq.Password = strings.TrimSpace(userReq.Password)

	if userReq.Username == "" || userReq.Password == "" {
		resError.WriteMessage(http.StatusBadRequest, "missing required fields")
		return resError
	}

	token, err := u.userRepo.SignIn(userReq)
	if err != nil {
		resError.WriteMessage(http.StatusBadRequest, "invalid credentials")
		return resError
	}

	w.Header().Set("Authorization", "Bearer "+token)

	var resOK httpresponse.OKResponse
	resOK.WriteMessage(http.StatusOK, "User successfully logged in")
	return resOK
}

func (u *UserHttpHandler) SignOutHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	user, _ := account.GetCurrentUser(r)
	// Blacklist this token so it can't be reused
	jwt.Store.Revoke(user.ID, user.ExpiresAt.Time)

	var resOK httpresponse.OKResponse
	resOK.WriteMessage(http.StatusOK, "You've successfully logged out")
	return resOK
} 
