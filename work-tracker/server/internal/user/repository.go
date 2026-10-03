package user

type UserRepository interface {
	GetByUsername(username string) (*UsersEntity, error)
	Create(user userRequest) (string, error)
}