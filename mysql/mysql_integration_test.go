// +build integration

package mysql

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"techtalk/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db IMySQL

func TestMain(m *testing.M) {
	dbusername := "root"
	dbpassword := "password1!"
	dbname := "mysqldb"
	dbinfo := fmt.Sprintf("%s:%s@/%s", dbusername, dbpassword, dbname)
	conn, err := sql.Open("mysql", dbinfo)
	if err != nil {
		panic(err)
	}
	db, err = NewSMySQL(conn)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestGetChannelMessage(t *testing.T) {
	randomID := rand.Intn(100)
	messages, err := db.GetChannelMessage(strconv.Itoa(randomID))
	if err != nil {
		t.Fatal(err)
	}
	if len(messages) < 100 {
		t.Fatal("result row is too little")
	}
}

func TestCreateChannelMessage(t *testing.T) {
	me := models.Message{
		ChannelID: "1",
		UserID:    "1",
		Type:      "test-type",
		Body:      "test-body",
		CreatedAt: 1561939200,
		UpdatedAt: 1561939200,
	}

	err := db.CreateChannelMessage(&me)
	if err != nil {
		t.Fatal(err)
	}
	if me.ID == "" {
		t.Fatal("new message may not be inserted")
	}
}

func TestUpdateChannelMessage(t *testing.T) {
	me := models.Message{
		ChannelID: "1",
		UserID:    "1",
		Type:      "test type",
		Body:      "test body",
		CreatedAt: 1561939200,
		UpdatedAt: 1561939200,
	}
	err := db.CreateChannelMessage(&me)
	if err != nil {
		t.Fatal(err)
	}
	if me.ID == "" {
		t.Fatal("new message may not be inserted")
	}

	me.Type = "updated body"
	me.Body = "updated body"
	me.CreatedAt = 1561939201
	me.UpdatedAt = 1561939202

	err = db.UpdateChannelMessage(&me)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteChannelMessage(t *testing.T) {
	me := models.Message{
		ChannelID: "1",
		UserID:    "1",
		Type:      "test type",
		Body:      "test body",
		CreatedAt: 1561939200,
		UpdatedAt: 1561939200,
	}
	err := db.CreateChannelMessage(&me)
	if err != nil {
		t.Fatal(err)
	}
	if me.ID == "" {
		t.Fatal("new message may not be inserted")
	}

	err = db.DeleteChannelMessage(me.ID)
	if err != nil {
		t.Fatal(err)
	}
}
