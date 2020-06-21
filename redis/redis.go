package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type IRedis interface {
	Ping() (string, error)
	Get(k string) (string, error)
	Set(k string, v string) error
}

type sRedis struct {
	rc *redis.Client
}

func NewSRedis(rc *redis.Client) (IRedis, error) {
	_, err := rc.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &sRedis{rc: rc}, nil
}

func (r *sRedis) Ping() (string, error) {
	return r.rc.Ping().Result()
}

func (r *sRedis) Get(key string) (string, error) {
	return r.rc.Get(key).Result()
}

func (r *sRedis) Set(key string, value string) error {
	return r.rc.Set(key, value, 60*time.Minute).Err()
}
