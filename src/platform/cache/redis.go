package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedisConnection(redisConnString string, ctxPool context.Context) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: redisConnString, // redisConnStr from .env » host:port
	})
	// ping redis for check connection
	_, err = rdb.Ping(ctxPool).Result()
	return rdb, err
}
