package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"techtalk/models"
	"techtalk/mysql"
	"techtalk/redis"
	"techtalk/utils"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetMessages(t *testing.T) {
	conn, err := utils.ConnectToDB()
	if err != nil {
		t.Fatal(err)
	}

	rc := utils.ConnectToRedis()

	db, err := mysql.NewSMySQL(conn)
	if err != nil {
		t.Fatal(err)
	}
	rds, err := redis.NewSRedis(rc)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	recorder := httptest.NewRecorder()
	attachHandlers(router, db, rds)
	req := httptest.NewRequest(http.MethodGet, "/channels/1/messages", nil)
	router.ServeHTTP(recorder, req)
	// fmt.Println(recorder.Body)
	ms := make([]models.Message, 0)
	err = json.Unmarshal(recorder.Body.Bytes(), &ms)
	if err != nil {
		t.Fatal(err)
	}
	if len(ms) < 100 {
		t.Fatal("result row is too little")
	}
}
