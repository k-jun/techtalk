package main

import (
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
