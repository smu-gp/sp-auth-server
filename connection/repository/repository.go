package repository

type ConnectionRepository interface {
	StoreSecret(userId string, secret string) (*string, error)
	GetSecret(key string) (*string, error)
	GetAllConnection() ([]string, error)
}
