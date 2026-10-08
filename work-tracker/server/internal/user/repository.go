package user

type UserRepository interface {
	GetByUsername(username string) (*UsersEntity, error)
	Create(user userSignUpRequest) (string, error)
	SignIn(user userSignInRequest) (string, error)
}