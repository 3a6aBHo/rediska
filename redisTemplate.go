package redislib

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func NewRedisClient(redisURL string, password string, database int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: password, // no password set
		DB:       database, // use default DB
	})
}

// Ping tests a connection to the Redis server.
func (c *Redis) Ping() (string, error) {
	return c.Client.Ping().Result()
}

// Get gets the value of a key.
func (c *Redis) Get(key string) (string, error) {
	return c.Client.Get(key).Result()
}

// Set sets the value of a key.
func (c *Redis) Set(key string, value string) error {
	return c.Client.Set(key, value, 0).Err()
}

// Del deletes one or more keys.
func (c *Redis) Del(keys ...string) error {
	return c.Client.Del(keys...).Err()
}

// Exists determines if a key exists.
func (c *Redis) Exists(key string) (int64, error) {
	return c.Client.Exists(key).Result()
}

// Expire sets a timeout on a key.
func (c *Redis) Expire(key string, duration int64) error {
	return c.Client.Expire(key, time.Duration(duration)).Err()
}

// TTL gets the remaining time to live of a key.
func (c *Redis) TTL(key string) (time.Duration, error) {
	return c.Client.TTL(key).Result()
}

// Incr increments the value of a key.
func (c *Redis) Incr(key string) (int64, error) {
	return c.Client.Incr(key).Result()
}

// Decr decrements the value of a key.
func (c *Redis) Decr(key string) (int64, error) {
	return c.Client.Decr(key).Result()
}

// HGet gets the value of a field in a hash.
func (c *Redis) HGet(key string, field string) (string, error) {
	return c.Client.HGet(key, field).Result()
}

// HSet sets the value of a field in a hash.
func (c *Redis) HSet(key string, field string, value string) error {
	return c.Client.HSet(key, field, value).Err()
}
