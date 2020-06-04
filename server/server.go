package server

import (
	"database/sql"
	"net/http"
	"techtalk/controllers"
	"techtalk/mysql"
	rds "techtalk/redis"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func NewServer(db *sql.DB, rc *redis.Client) http.Server {
	router := mux.NewRouter()

	mysqldb, err := mysql.NewSMySQL(db)
	if err != nil {
		panic(err)
	}

	rds, err := rds.NewSRedis(rc)
	if err != nil {
		panic(err)
	}
	attachHandlers(router, mysqldb, rds)

	s := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	return s
}

func attachHandlers(mux *mux.Router, db mysql.IMySQL, rds rds.IRedis) {
	mux.HandleFunc("/channels/{id}/messages", controllers.GetMessages(db, rds)).Methods(http.MethodGet)
	mux.HandleFunc("/channels/{id}/messages", controllers.PostMessage(db, rds)).Methods(http.MethodPost)
	mux.HandleFunc("/messages", controllers.PutMessage(db, rds)).Methods(http.MethodPut)
	mux.HandleFunc("/messages", controllers.DeleteMessage(db, rds)).Methods(http.MethodDelete)
}
