package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	pageTop = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Soundex</title>
	</head>
	<body>
		<h1>Soundex</h1>
		<p>Compute soundex codes for a list of names.</p>`
	form = `<form action="">
	<p><label for="names">Names (comma or space-separated)</label></p>
	<p><input type="text" name="names"></p>
	<p><input type="submit" value="Compute"></p>
</form>`
	pageBottom = `</body>
	</html>`
	anError = `<p class="error">%s</p>`
)

//viewResult handles requests to /test-page.
func viewResult(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

//mainForm handles requests to main page.
func mainForm(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, pageTop, form, pageBottom)
}

func main() {
	http.HandleFunc("/", mainForm)
	http.HandleFunc("/test", viewResult)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
