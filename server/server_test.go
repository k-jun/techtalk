package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"techtalk/models"
	"techtalk/mysql"
	"techtalk/redis"
	"techtalk/utils"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
)

var db mysql.IMySQL
var rc redis.IRedis

func TestMain(m *testing.M) {
	dbhost := "localhost"
	dbusername := "root"
	dbpassword := "password1!"
	dbname := "mysqldb"
	redisEndpoint := "localhost:6379"

	conn, err := utils.ConnectToDB(dbusername, dbpassword, dbhost, dbname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	rrc := utils.ConnectToRedis(redisEndpoint)
	defer rrc.Close()

	db, err = mysql.NewSMySQL(conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rc, err = redis.NewSRedis(rrc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestGetMessages(t *testing.T) {
	tests := []struct {
		url          string
		isBlankArray bool
	}{
		{url: "/channels/1/messages", isBlankArray: false},
		{url: "/channels/0/messages", isBlankArray: false},
		{url: "/channels/some_random_string/messages", isBlankArray: true},
	}

	for _, test := range tests {

		router := mux.NewRouter()
		recorder := httptest.NewRecorder()
		attachHandlers(router, db, rc)
		req := httptest.NewRequest(http.MethodGet, test.url, nil)
		router.ServeHTTP(recorder, req)

		ms := make([]models.Message, 0)
		err := json.Unmarshal(recorder.Body.Bytes(), &ms)
		if err != nil {
			t.Fatal(err)
		}
		if len(ms) == 0 != test.isBlankArray {
			t.Fatal("returned message number is invalid")
		}
	}
}

func TestPostMessages(t *testing.T) {

	tests := []struct {
		url  string
		body string
		code int
	}{
		{
			url:  "/channels/1/messages",
			body: `{"user_id": "1", "type": "%s", "body": "%s"}`,
			code: 200,
		},
		{
			url:  "/channels/0/messages",
			body: `{"user_id": "1", "type": "%s", "body": "%s"}`,
			code: 200,
		},
		{
			url:  "/channels/1/messages",
			body: `{"user_id": "1", "type": "%s", "body": "%s"`,
			code: 400,
		},
		{
			url:  "/channels/some_random_string/messages",
			body: `{"user_id": "1", "type": "%s", "body": "%s"}`,
			code: 400,
		},
	}

	for _, test := range tests {

		randomType := RandStringRunes(10)
		randomBody := RandStringRunes(10)

		router := mux.NewRouter()
		recorder := httptest.NewRecorder()
		attachHandlers(router, db, rc)
		body := fmt.Sprintf(test.body, randomType, randomBody)
		bodyByte := []byte(body)
		req := httptest.NewRequest(http.MethodPost, test.url, bytes.NewBuffer(bodyByte))
		router.ServeHTTP(recorder, req)
		assert.Equal(t, recorder.Code, test.code)

		if recorder.Code == http.StatusOK {
			var m models.Message
			err := json.Unmarshal(recorder.Body.Bytes(), &m)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, m.Type, randomType)
			assert.Equal(t, m.Body, randomBody)
		}

	}
}

func TestPutMessages(t *testing.T) {

	tests := []struct {
		url        string
		createBody string
		updateBody string
		code       int
	}{
		{
			url:        "/channels/1/messages",
			createBody: `{"user_id": "1", "type": "sample type", "body": "sample body"}`,
			updateBody: `{"id": "%s", "type": "%s", "body": "%s"}`,
			code:       200,
		},
		{
			url:        "/channels/1/messages",
			createBody: `{"user_id": "1", "type": "sample type", "body": "sample body"}`,
			updateBody: `{"id": "%s", "type": "%s", "body": "%s"`,
			code:       400,
		},
		{
			url:        "/channels/1/messages",
			createBody: `{"user_id": "1", "type": "sample type", "body": "sample body"}`,
			updateBody: `{"id": "0", "type": "%s%s", "body": "%s"}`,
			code:       400,
		},
	}

	for _, test := range tests {
		randomType := RandStringRunes(10)
		randomBody := RandStringRunes(10)

		router := mux.NewRouter()
		attachHandlers(router, db, rc)

		// make create request
		recorder := httptest.NewRecorder()
		bodyByte := []byte(test.createBody)
		req := httptest.NewRequest(http.MethodPost, test.url, bytes.NewBuffer(bodyByte))
		router.ServeHTTP(recorder, req)

		// check result
		if recorder.Code != http.StatusOK {
			t.Fatal("createBody is invalid")
		}
		var m models.Message
		err := json.Unmarshal(recorder.Body.Bytes(), &m)
		if err != nil {
			t.Fatal(err)
		}

		// make update request
		recorder = httptest.NewRecorder()
		updateBody := fmt.Sprintf(test.updateBody, m.ID, randomType, randomBody)
		bodyByte = []byte(updateBody)
		req = httptest.NewRequest(http.MethodPut, test.url, bytes.NewBuffer(bodyByte))
		router.ServeHTTP(recorder, req)

		// check result
		assert.Equal(t, recorder.Code, test.code)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestDeleteMessages(t *testing.T) {

	tests := []struct {
		url        string
		createBody string
		deleteBody string
		code       int
	}{
		{
			url:        "/channels/1/messages",
			createBody: `{"user_id": "1", "type": "sample type", "body": "sample body"}`,
			deleteBody: `{"id": "%s"}`,
			code:       200,
		},
		{
			url:        "/channels/1/messages",
			createBody: `{"user_id": "1", "type": "sample type", "body": "sample body"}`,
			deleteBody: `{"id": "%s999"`,
			code:       400,
		},
		{
			url:        "/channels/1/messages",
			createBody: `{"user_id": "1", "type": "sample type", "body": "sample body"}`,
			deleteBody: `{"id": "%s999"}`,
			code:       400,
		},
	}

	for _, test := range tests {

		router := mux.NewRouter()
		attachHandlers(router, db, rc)

		// make create request
		recorder := httptest.NewRecorder()
		bodyByte := []byte(test.createBody)
		req := httptest.NewRequest(http.MethodPost, test.url, bytes.NewBuffer(bodyByte))
		router.ServeHTTP(recorder, req)

		// check result
		if recorder.Code != http.StatusOK {
			t.Fatal("createBody is invalid")
		}
		var m models.Message
		err := json.Unmarshal(recorder.Body.Bytes(), &m)
		if err != nil {
			t.Fatal(err)
		}

		// make update request
		recorder = httptest.NewRecorder()
		deleteBody := fmt.Sprintf(test.deleteBody, m.ID)
		bodyByte = []byte(deleteBody)
		req = httptest.NewRequest(http.MethodDelete, test.url, bytes.NewBuffer(bodyByte))
		router.ServeHTTP(recorder, req)

		// check result
		assert.Equal(t, recorder.Code, test.code)
	}
}
