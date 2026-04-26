package repository

type AddressRepository interface {
	Create(addr map[string]any, userID int32) (*addressAggre, error)
	GetAddresses(userID, limit, offset int32) (*[]addressAggreTtl, error)
	Delete(addrId int32) (int32, error)
	SetDefault(addrID int32, userID int32) (int32, error)
	//Update(addrID map[string]any, id int32) (*addressAggre, error)
}
