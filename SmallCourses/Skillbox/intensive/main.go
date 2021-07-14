package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/goombaio/namegenerator"
)

//hello outputs random string in browser.
func hello(w http.ResponseWriter, r *http.Request) {
	generator := namegenerator.NewNameGenerator(rand.Int63())
	fmt.Fprintf(w, generator.Generate())
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8001", nil)
}
