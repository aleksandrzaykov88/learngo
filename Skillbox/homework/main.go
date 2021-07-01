package main

import (
	"fmt"
	"log"
	"net/http"
)

//hello process input URL and greetings this URL-name.
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
