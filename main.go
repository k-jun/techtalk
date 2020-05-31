package main

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"techtalk/controllers"
)

var db *sql.DB
var rc *redis.Client

const (
	RedisAddress = "localhost:6379"
)

func main() {
	http.Handle("/messages", http.HandlerFunc(controllers.GetMessages))

	http.ListenAndServe(":8000", nil)

}
