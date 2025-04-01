package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Storage struct {
	client *redis.Client
}

func New(address string) *Storage {
	return &Storage{client: redis.NewClient(&redis.Options{
		Addr: address,
	})}
}

func (s *Storage) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return s.client.Set(
		ctx, key, value, ttl,
	).Err()
}

func (s *Storage) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := s.client.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (s *Storage) Get(ctx context.Context, key string) ([]byte, error) {
	return s.client.Get(ctx, key).Bytes()
}
