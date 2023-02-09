package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisTemplate struct {
	Client *redis.Client
}

func NewRedisClient(url string, password string, db int) *RedisTemplate {
	return &RedisTemplate{
		Client: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: password,
			DB:       db,
		}),
	}
}

func (s *RedisTemplate) Set(key string, value string) error {
	err := s.Client.Set(key, value, 0).Err()
	return err
}

func (s *RedisTemplate) Get(key string) (string, error) {
	val, err := s.Client.Get(key).Result()
	return val, err
}

// Del deletes one or more keys.
func (s *RedisTemplate) Delete(key ...string) error {
	err := s.Client.Del(key...).Err()
	return err
}

// Exists determines if a key exists.
func (s *RedisTemplate) Exists(key string) (int64, error) {
	return s.Client.Exists(key).Result()
}

// Expire sets a timeout on a key.
func (s *RedisTemplate) Expire(key string, duration int64) error {
	return s.Client.Expire(key, time.Duration(duration)).Err()
}

// TTL gets the remaining time to live of a key.
func (s *RedisTemplate) TTL(key string) (time.Duration, error) {
	return s.Client.TTL(key).Result()
}

// Incr increments the value of a key.
func (s *RedisTemplate) Incr(key string) (int64, error) {
	return s.Client.Incr(key).Result()
}

// Decr decrements the value of a key.
func (s *RedisTemplate) Decr(key string) (int64, error) {
	return s.Client.Decr(key).Result()
}

// HGet gets the value of a field in a hash.
func (s *RedisTemplate) HGet(key string, field string) (string, error) {
	return s.Client.HGet(key, field).Result()
}

// HSet sets the value of a field in a hash.
func (s *RedisTemplate) HSet(key string, field string, value string) error {
	return s.Client.HSet(key, field, value).Err()
}

//ping function
func (s *RedisTemplate) Ping() (string, error) {
	return s.Client.Ping().Result()
}

//close connection
func (s *RedisTemplate) CloseConnection() error {
	err := s.Client.Close()
	return err
}
