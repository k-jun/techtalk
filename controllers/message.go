package controllers

import (
	"encoding/json"
	"net/http"
	"techtalk/mysql"
	"techtalk/redis"

	"github.com/gorilla/mux"
)

func GetMessages(db mysql.IMySQL, rds redis.IRedis) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		messages, err := db.GetChannelMessage(vars["id"])
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
