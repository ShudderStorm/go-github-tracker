package cache

import (
	"context"
	"time"
)

type KeyValStorage interface {
	Set(context.Context, string, string, time.Duration) error
	Exists(context.Context, string) (bool, error)
	Get(context.Context, string) (string, error)
}
