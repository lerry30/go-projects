package user

type userSignUpRequest struct {
	FirstName string `json:"first-name"`
	LastName string `json:"last-name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type userSignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}