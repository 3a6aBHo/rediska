package rediska

import (
	"github.com/go-redis/redis"
	"time"
)

// Client is a custom Redis client.
type Client struct {
	RedisClient *redis.Client
}

// NewClient creates a new Redis client.
func NewClient(options *redis.Options) (*Client, error) {
	redisClient := redis.NewClient(options)
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &Client{RedisClient: redisClient}, nil
}

// Ping tests a connection to the Redis server.
func (c *Client) Ping() (string, error) {
	return c.RedisClient.Ping().Result()
}

// Get gets the value of a key.
func (c *Client) Get(key string) (string, error) {
	return c.RedisClient.Get(key).Result()
}

// Set sets the value of a key.
func (c *Client) Set(key string, value string) error {
	return c.RedisClient.Set(key, value, 0).Err()
}

// Del deletes one or more keys.
func (c *Client) Del(keys ...string) error {
	return c.RedisClient.Del(keys...).Err()
}

// Exists determines if a key exists.
func (c *Client) Exists(key string) (int64, error) {
	return c.RedisClient.Exists(key).Result()
}

// Expire sets a timeout on a key.
func (c *Client) Expire(key string, duration int64) error {
	return c.RedisClient.Expire(key, time.Duration(duration)).Err()
}

// TTL gets the remaining time to live of a key.
func (c *Client) TTL(key string) (time.Duration, error) {
	return c.RedisClient.TTL(key).Result()
}

// Incr increments the value of a key.
func (c *Client) Incr(key string) (int64, error) {
	return c.RedisClient.Incr(key).Result()
}

// Decr decrements the value of a key.
func (c *Client) Decr(key string) (int64, error) {
	return c.RedisClient.Decr(key).Result()
}

// HGet gets the value of a field in a hash.
func (c *Client) HGet(key string, field string) (string, error) {
	return c.RedisClient.HGet(key, field).Result()
}

// HSet sets the value of a field in a hash.
func (c *Client) HSet(key string, field string, value string) error {
	return c.RedisClient.HSet(key, field, value).Err()
}

// HGetAll gets all fields and values in a hash.
func (c *Client) HGetAll(key string, field string) (string, error) {
	return c.RedisClient.HGet(key, field).Result()

}
