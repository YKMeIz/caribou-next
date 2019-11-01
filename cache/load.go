package cache

import (
	"errors"
	"github.com/go-redis/redis"
)

var NoKeyError = errors.New("redis: key does not exist")

func LoadBytes(k string) ([]byte, error) {
	b, err := client.Get(k).Bytes()
	if err == redis.Nil {
		return nil, NoKeyError
	}

	return b, err
}

func LoadString(k string) (string, error) {
	s, err := client.Get(k).Result()
	if err == redis.Nil {
		return "", NoKeyError
	}

	return s, err
}
