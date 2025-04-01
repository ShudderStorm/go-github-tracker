package oauth

import (
	"context"
	"github.com/ShudderStorm/go-github-tracker/internal/redis"
	"time"
)

type RedisStateHolder struct {
	storage *redis.Storage
}

func NewRedisStateHolder(storage *redis.Storage) *RedisStateHolder {
	return &RedisStateHolder{storage: storage}
}

func (r RedisStateHolder) Store(state string, ttl time.Duration) error {
	return r.storage.Set(
		context.Background(),
		state,
		[]byte{255},
		ttl,
	)
}

func (r RedisStateHolder) Validate(state string) (bool, error) {
	return r.storage.Exists(context.Background(), state)
}
