package repository

import "github.com/go-redis/redis"

type redisConnectionRepository struct {
	client *redis.Client
}

func NewRedisConnectionRepository(client *redis.Client) ConnectionRepository {
	return &redisConnectionRepository{client}
}

func (redisConnectionRepository) StoreCode(userId string, code int32) (*string, error) {
	panic("implement me")
}

func (redisConnectionRepository) GetUserIdByCode(code int32) (*string, error) {
	panic("implement me")
}
