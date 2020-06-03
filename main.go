package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"techtalk/server"
	"techtalk/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	conn, err := utils.ConnectToDB()
	rc := utils.ConnectToRedis()

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
