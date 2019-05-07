package repository

type ConnectionRepository interface {
	StoreSecret(userId string, secret string) (*string, error)
	GetSecret(key string) (*string, error)
	DeleteKey(key string) (bool, error)
	GetAllConnection() ([]string, error)
}
