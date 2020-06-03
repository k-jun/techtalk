package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"techtalk/mysql"
	"techtalk/redis"

	"github.com/gorilla/mux"
)

func GetMessages(db mysql.IMySQL, rds redis.IRedis) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cid := retrieveId(r)
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

func retrieveId(r *http.Request) string {
	// if number is 0, generate random number
	vars := mux.Vars(r)
	cid := vars["id"]
	if cid == "0" {
		cid = strconv.Itoa(rand.Intn(1000))
	}

	return cid
}
