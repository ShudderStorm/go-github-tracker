package cache

import (
	"context"
	"fmt"
	"time"
)

type KeyValStorage interface {
	Set(context.Context, string, []byte, time.Duration) error
	Exists(context.Context, string) (bool, error)
	Get(context.Context, string) ([]byte, error)
}

type Cache[V any] struct {
	storage KeyValStorage
	ttl     time.Duration

	serialize   func(V) ([]byte, error)
	deserialize func([]byte) (V, error)
}

func New[V any](storage KeyValStorage, serialize func(V) ([]byte, error), deserialize func([]byte) (V, error)) *Cache[V] {
	return &Cache[V]{
		storage:     storage,
		serialize:   serialize,
		deserialize: deserialize,
	}
}

func (c *Cache[V]) WithTTL(ttl time.Duration) *Cache[V] {
	c.ttl = ttl
	return c
}

func (c *Cache[V]) Store(ctx context.Context, key string, value V) error {
	bytes, err := c.serialize(value)

	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}

	return c.storage.Set(ctx, key, bytes, c.ttl)
}

func (c *Cache[V]) Valid(ctx context.Context, key string) (bool, error) {
	return c.storage.Exists(ctx, key)
}

func (c *Cache[V]) Fetch(ctx context.Context, key string) (V, error) {
	var value V
	bytes, err := c.storage.Get(ctx, key)
	if err != nil {
		return value, err
	}

	value, err = c.deserialize(bytes)
	if err != nil {
		return value, fmt.Errorf("deserialization error: %w", err)
	}

	return value, nil
}
