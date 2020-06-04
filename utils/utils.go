package utils

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(dbusername, dbpassword, dbhost, dbname string) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbusername, dbpassword, dbhost, dbname)
	conn, err := sql.Open("mysql", dbinfo)
	return conn, err
}

func ConnectToRedis(endpoint string) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: endpoint})
}
