package router

import (
	"fmt"
	"net/http"
)

type DBConn interface {
	Ping() error
}

func New(db DBConn) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingDB(db))
	return mux
}

func pingDB(db DBConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(
				w, fmt.Sprintf("db ping error: %v", err), http.StatusInternalServerError,
			)
			return
		}
		w.Write([]byte("pong!"))
	}
}
