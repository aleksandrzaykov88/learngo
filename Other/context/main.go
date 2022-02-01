package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type key int

const (
	userIDctx int = 0
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Id")

	ctx := context.WithValue(r.Context(), userIDctx, id)

	result := processLongTask(ctx)

	w.Write([]byte(result))
}

func processLongTask(ctx context.Context) string {
	id := ctx.Value(userIDctx)
	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintln("done processing id", id)
	case <-ctx.Done():
		log.Println("request canceled")
		return ""
	}
}
