package redislib

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}
type OptionsRedis struct {
	Options *redis.Options
}

func NewConnection(addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Ping tests a connection to the Redis server.
func (c *RedisClient) Ping() (string, error) {
	return c.Client.Ping().Result()
}

// Get gets the value of a key.
func (c *RedisClient) Get(key string) (string, error) {
	return c.Client.Get(key).Result()
}

// Set sets the value of a key.
func (c *RedisClient) Set(key string, value string) error {
	return c.Client.Set(key, value, 0).Err()
}

// Del deletes one or more keys.
func (c *RedisClient) Del(keys ...string) error {
	return c.Client.Del(keys...).Err()
}

// Exists determines if a key exists.
func (c *RedisClient) Exists(key string) (int64, error) {
	return c.Client.Exists(key).Result()
}

// Expire sets a timeout on a key.
func (c *RedisClient) Expire(key string, duration int64) error {
	return c.Client.Expire(key, time.Duration(duration)).Err()
}

// TTL gets the remaining time to live of a key.
func (c *RedisClient) TTL(key string) (time.Duration, error) {
	return c.Client.TTL(key).Result()
}

// Incr increments the value of a key.
func (c *RedisClient) Incr(key string) (int64, error) {
	return c.Client.Incr(key).Result()
}

// Decr decrements the value of a key.
func (c *RedisClient) Decr(key string) (int64, error) {
	return c.Client.Decr(key).Result()
}

// HGet gets the value of a field in a hash.
func (c *RedisClient) HGet(key string, field string) (string, error) {
	return c.Client.HGet(key, field).Result()
}

// HSet sets the value of a field in a hash.
func (c *RedisClient) HSet(key string, field string, value string) error {
	return c.Client.HSet(key, field, value).Err()
}

// HGetAll gets all fields and values in a hash.
func (c *RedisClient) HGetAll(key string, field string) (string, error) {
	return c.Client.HGet(key, field).Result()

}
