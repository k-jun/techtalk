package server

import (
	"context"
	"database/sql"
	"net/http"
	"techtalk/controllers"
	"techtalk/mysql"
	rds "techtalk/redis"
	"time"

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
	attachHandlersWithContext(router, mysqldb, rds)

	s := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	return s
}

func attachHandlers(mux *mux.Router, db mysql.IMySQL, rds rds.IRedis) {
	mux.HandleFunc("/channels/{id}/messages", controllers.GetMessages(db, rds)).Methods(http.MethodGet)
}

func attachHandlersWithContext(mux *mux.Router, db mysql.IMySQL, rds rds.IRedis) {
	muxTimeout := mux.PathPrefix("").Subrouter()
	muxTimeout.Use(WithTimeout)
	muxTimeout.HandleFunc("/channels/{id}/messages", controllers.PostMessage(db, rds)).Methods(http.MethodPost)
	muxTimeout.HandleFunc("/channels/{id}/messages", controllers.PutMessage(db, rds)).Methods(http.MethodPut)
	muxTimeout.HandleFunc("/channels/{id}/messages", controllers.DeleteMessage(db, rds)).Methods(http.MethodDelete)
}

func WithTimeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		ch := make(chan string)

		go func() {
			next.ServeHTTP(w, r)
			ch <- "done"
		}()
		select {
		case <-ch:
			// complete before timeout
		case <-ctx.Done():
			// not complete within 100ms
		}
		return
	})
}
