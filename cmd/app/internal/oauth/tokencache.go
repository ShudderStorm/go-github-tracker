package oauth

import (
	"context"
	"github.com/ShudderStorm/go-github-tracker/internal/redis"
	"strconv"
	"time"
)

const DefaultTokenTTl time.Duration = time.Hour

type TokenCache struct {
	storage *redis.Storage
}

func NewTokenCache(storage *redis.Storage) *TokenCache {
	return &TokenCache{storage: storage}
}

func (cache *TokenCache) Store(id int, token string, ttl time.Duration) error {
	return cache.storage.Set(
		context.Background(),
		strconv.Itoa(id),
		[]byte(token),
		ttl,
	)
}

func (cache *TokenCache) Get(id int) (string, error) {
	tokenbytes, err := cache.storage.Get(context.Background(), strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return string(tokenbytes), nil
}
