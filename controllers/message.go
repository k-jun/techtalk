package controllers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"techtalk/models"
	"techtalk/mysql"
	"techtalk/redis"

	"github.com/gorilla/mux"
)

func GetMessages(db mysql.IMySQL, rds redis.IRedis) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cid := retrieveIdFromPath(r)
		messages, err := db.GetChannelMessage(cid)
		if err != nil {
			BadRequest(w, r)
			return
		}

		bytes, err := json.Marshal(messages)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		w.Write(bytes)
	}
}

func PostMessages(db mysql.IMySQL, rds redis.IRedis) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cid := retrieveIdFromPath(r)
		m, err := retrieveMessageFromBody(r)
		if err != nil {
			BadRequest(w, r)
			return
		}
		m.ChannelID = cid
		err = db.CreateChannelMessage(&m)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		bytes, err := json.Marshal(m)
		w.Write(bytes)

	}
}

func PutMessages(db mysql.IMySQL, rds redis.IRedis) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := retrieveMessageFromBody(r)
		if err != nil {
			BadRequest(w, r)
			return
		}
		err = db.UpdateChannelMessage(&m)
		if err != nil {
			InternalServerError(w, r)
			return
		}
	}
}

// func DeleteMessage(db mysql.IMySQL, rds redis.IRedis) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m, err := retrieveMessageFromBody(r)
// 		if err != nil {
// 			BadRequest(w, r)
// 			return
// 		}
// 		err = db.DeleteChannelMessage(m.ID)
// 		if err != nil {
// 			fmt.Println(err)
// 			InternalServerError(w, r)
// 			return
// 		}
// 	}
// }

func retrieveIdFromPath(r *http.Request) string {
	// if number is 0, generate random number
	vars := mux.Vars(r)
	cid := vars["id"]
	if cid == "0" {
		cid = strconv.Itoa(rand.Intn(1000) + 1)
	}

	return cid
}

func retrieveMessageFromBody(r *http.Request) (models.Message, error) {
	var m models.Message

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(bytes, &m)
	return m, err
}
