// +build integration

package redis

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
)

var r IRedis

func TestMain(m *testing.M) {
	var err error

	rc := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	r, err = NewSRedis(rc)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestRedisPing(t *testing.T) {
	pong, err := r.Ping()
	if err != nil {
		t.Fatal(err)
	}
	if pong != "PONG" {

		t.Fatal("message was not PONG")
	}
}
func TestRedisGet(t *testing.T) {
	_, err := r.Get("non-value")
	if err != nil && err != redis.Nil {
		t.Fatal(err)
	}
}

func TestRedisSet(t *testing.T) {
	err := r.Set("sample-key", "sample-value")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedisSetAndGet(t *testing.T) {
	key := "qawsedrf"
	value := "rdftgyhujikolp"

	err := r.Set(key, value)
	if err != nil {
		t.Fatal(err)
	}
	got, err := r.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if got != value {
		t.Fatalf("value is not same. got: %s", got)
	}
}
