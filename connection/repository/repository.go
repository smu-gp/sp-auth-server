package repository

type ConnectionRepository interface {

	StoreCode(userId string, code int32) (*string, error)

	GetUserIdByCode(code int32) (*string, error)

}