package redis_client

import "github.com/go-redis/redis/v8"

var (
	ErrNil    = redis.Nil
	ErrClosed = redis.ErrClosed
)