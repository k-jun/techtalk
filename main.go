package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"os"
	"techtalk/server"
	"techtalk/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbhost := os.Getenv("DB_HOST")
	dbusername := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	redisEndpoint := os.Getenv("REDIS_ENDPOINT")

	conn, err := utils.ConnectToDB(dbusername, dbpassword, dbhost, dbname)
	rc := utils.ConnectToRedis(redisEndpoint)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	s := server.NewServer(conn, rc)
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func init() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}
