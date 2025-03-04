package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Serializable interface {
	Serialize() (string, error)
}

type Deserializer func(string) (Serializable, error)

type Storage struct {
	client *redis.Client
}

func New(host string, port uint16) *Storage {
	return &Storage{client: redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})}
}

func (s *Storage) Set(ctx context.Context, key string, value Serializable, ttl time.Duration) error {
	serialized, err := value.Serialize()
	if err != nil {
		return fmt.Errorf(
			"failed to serialize value: %w", err,
		)
	}

	return s.client.Set(
		ctx, key, serialized, ttl,
	).Err()
}

func (s *Storage) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := s.client.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (s *Storage) Get(ctx context.Context, key string, deserializer Deserializer) (Serializable, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	res, err := deserializer(val)
	if err != nil {
		return nil, fmt.Errorf(
			"deserializer failed: %w", err,
		)
	}

	return res, nil
}
