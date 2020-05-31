// +build integration

package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var m IMySQL

func TestMain(t *testing.M) {
	dbinfo := fmt.Sprintf("%s:%s@/%s", dbusername, dbpassword, dbname)
	conn, err := sql.Open(dbinfo)
}
