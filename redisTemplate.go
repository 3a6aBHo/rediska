package redislib

import (
	"github.com/go-redis/redis"
)

func NewRedisClient(redisURL string, password string, database int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: password, // no password set
		DB:       database, // use default DB
	})
}
