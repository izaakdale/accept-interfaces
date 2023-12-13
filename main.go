package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/izaakdale/accept-interfaces/db"
	"github.com/izaakdale/accept-interfaces/router"
)

func main() {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	cli := redis.NewClient(opts)
	conn, err := db.New(cli)
	if err != nil {
		log.Fatal(err)
	}

	mux := router.New(conn)
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		Handler: mux,
	}
	log.Printf("listening on %s...\n", srv.Addr)
	srv.ListenAndServe()
}
