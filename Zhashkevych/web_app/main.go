package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/aleksandrzaykov88/learngo/Zhashkevych/geometric_shapes/circle"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		t.ExecuteTemplate(w, "index", "")
	})
	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		radius := r.URL.Query().Get("radius")
		var circle circle.Circle
		number, err := strconv.ParseFloat(radius, 64)
		if err != nil {
			return
		}
		circle.R = number
		area := circle.GetArea()
		fmt.Fprint(w, "Area of circle with radius ", radius, " is equal ", area)
	})
	http.ListenAndServe(":9001", nil)
}
