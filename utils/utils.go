package utils

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	dbusername := "root"
	dbpassword := "password1!"
	dbname := "mysqldb"
	dbinfo := fmt.Sprintf("%s:%s@/%s", dbusername, dbpassword, dbname)
	conn, err := sql.Open("mysql", dbinfo)
	return conn, err
}

func ConnectToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379"})
}
