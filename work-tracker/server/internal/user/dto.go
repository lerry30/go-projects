package user

type userRequest struct {
	FirstName string `json:"first-name"`
	LastName string `json:"last-name"`
	Username string `json:"username"`
	Password string `json:"password"`
}