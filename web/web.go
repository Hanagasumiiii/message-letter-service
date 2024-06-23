package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		http.Error(w, "Please provide a message", http.StatusBadRequest)
	}
	err := rdb.LPush(ctx, "message", message).Err()
	if err != nil {
		log.Fatal(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = fmt.Fprintf(w, "Message: %s sent", message)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	})
}

func main() {
	http.HandleFunc("/notify/", messageHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
