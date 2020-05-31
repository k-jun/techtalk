package redis

import (
	"os"

	"github.com/go-redis/redis"
)

var rc *redis.Client

func init() {
	re := os.Getenv("REDIS_ENDPOINT")
	if re == "" {
		panic("environment valiable was not set")
	}

	rc = redis.NewClient(&redis.Options{Addr: re})
	_, err := rc.Ping().Result()
	if err != nil {
		panic(err)
	}
}
