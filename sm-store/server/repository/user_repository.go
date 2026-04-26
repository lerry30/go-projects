package repository

type UserRepository interface {
	GetByID(id int32) (*userRow, error)
	GetByEmail(email string) (*userRow, error)
	Create(user map[string]any) (int32, error)
	Update(user map[string]any, id int32) (*userRow, error)
}
