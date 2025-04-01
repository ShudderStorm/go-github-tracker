package redis

import (
	"github.com/ShudderStorm/go-github-tracker/internal/redis"
	"os"
)

const RedisAddrEnvKey = "REDIS_ADDR"

var Storage *redis.Storage

func init() {
	Storage = redis.New(os.Getenv(RedisAddrEnvKey))
}
