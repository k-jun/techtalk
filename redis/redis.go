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

func (r *Redis) Ping() (string, error) {
	return r.rc.Ping().Result()
}

func (r *Redis) Get(key string) (string, error) {
	return r.rc.Get(key).Result()
}

func (r *Redis) Set(key string, value string) error {
	return r.rc.Set(key, value, 0).Err()
}
