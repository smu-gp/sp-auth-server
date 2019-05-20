package repository

import "github.com/go-redis/redis"

type ConnectionRepository interface {
	StoreSecret(userId string, secret string) (*string, error)
	GetSecret(key string) (*string, error)
	DeleteKey(key string) (bool, error)
	GetAllConnection() ([]string, error)
	Subscribe(channel string) *redis.PubSub
	Publish(channel string, message string) (int, error)
}
