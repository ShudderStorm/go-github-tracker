package cache

import (
	"context"
	"time"
)

type KeyValStorage interface {
	Set(context.Context, string, []byte, time.Duration) error
	Exists(context.Context, string) (bool, error)
	Get(context.Context, string) ([]byte, error)
}
