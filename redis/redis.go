package redis

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	rc *redis.Client
}

func NewRedisClient(a string) (*Redis, error) {
	rc := redis.NewClient(&redis.Options{Addr: a})
	_, err := rc.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &Redis{rc: rc}, nil
}

func Get(key string) (string, error) {
	// use redis client to get data by key

	return "", nil
}

func Set(key string, value string) error {
	// use redis client to store data by key
	return nil
}
