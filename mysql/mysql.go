package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	dbuser := os.Getenv("DB_USER")
	// dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("%s@/%s", dbuser, dbname)
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
