package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/goombaio/namegenerator"
)

func hello(w http.ResponseWriter, r *http.Request) {
	generator := namegenerator.NewNameGenerator(rand.Int63())
	fmt.Fprintf(w, "Hello %s", generator.Generate())
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8000", nil)
}
