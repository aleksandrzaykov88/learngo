package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"regexp"
	"strings"
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
	<body>`
	pageTitle = `<h1>Soundex</h1>
		<p>Compute soundex codes for a list of names.</p>`
	form = `<form  action="/test" method="POST">
	<p><label for="names">Names (comma or space-separated)</label></p>
	<p><input type="text" name="names"></p>
	<p><input type="submit" value="Compute"></p>
</form>`
	pageBottom = `</body>
	</html>`
)

//letterInSlice returns true if letter is in slice.
func letterInSlice(letter string, slice []string) bool {
	for _, l := range slice {
		if l == letter {
			return true
		}
	}
	return false
}

//letterToNum describes the third step of soundex algorithm.
//Certain letters are assigned special numbers.
func letterToNum(letter string) string {
	firstSet := []string{"b", "f", "p", "v"}
	secondSet := []string{"c", "g", "j", "k", "q", "s", "x", "z"}
	thirdSet := []string{"d", "t"}
	fourthSet := []string{"l"}
	fifthSet := []string{"m", "n"}
	sixthSet := []string{"r"}
	seventhSet := []string{"a", "e", "i", "o", "u", "y"}

	switch {
	case letterInSlice(letter, firstSet):
		return "1"
	case letterInSlice(letter, secondSet):
		return "2"
	case letterInSlice(letter, thirdSet):
		return "3"
	case letterInSlice(letter, fourthSet):
		return "4"
	case letterInSlice(letter, fifthSet):
		return "5"
	case letterInSlice(letter, sixthSet):
		return "6"
	case letterInSlice(letter, seventhSet):
		return "0"
	default:
		return ""
	}
}

//removeVovels describes the fifth step of soundex algorithm.
//All a, e, i, o, u, y -letters removes from word.
func removeVovels(word string) string {
	var r rune
	for _, c := range word {
		r = c
		break
	}
	newWord := string(r)
	for i := 1; i <= len(word)-1; i++ {
		if string(word[i]) == "0" {
			continue
		}
		newWord += string(word[i])
	}
	return newWord
}

//removeDoubles describes the fifth step of soundex algorithm.
//Any sequence of identical digits is reduced to one such digit.
func removeDoubles(word string) string {
	var r rune
	for _, c := range word {
		r = c
		break
	}
	newWord := string(r)
	for i := 1; i <= len(word)-1; i++ {
		if i == 1 {
			newWord += letterToNum(string(word[i]))
		} else if i > 1 {
			a := letterToNum(string(word[i-1]))
			b := letterToNum(string(word[i]))
			if b != a {
				newWord += b
			}
		}
	}
	return newWord
}

//soundex sets the same index for strings that sound similar in English.
func soundex(word string) string {
	var r rune
	for _, c := range word {
		r = c
		break
	}
	newWord := strings.ToUpper(string(r))
	for i := 1; i <= len(word)-1; i++ {
		if string(word[i]) == "h" || string(word[i]) == "w" {
			continue
		}
		newWord += string(word[i])
	}
	newWord = removeDoubles(newWord)
	newWord = removeVovels(newWord)
	for {
		if len(newWord) < 4 {
			newWord += "0"
		} else {
			break
		}
	}
	return newWord[:4]
}

//viewResult handles requests to /test-page.
func viewResult(writer http.ResponseWriter, request *http.Request) {
	splitter := regexp.MustCompile(`( *, *)|(  *)`)
	names := splitter.Split(request.FormValue("names"), -1)
	var newNames []string
	for _, name := range names {
		matched, err := regexp.MatchString(`\b([a-zA-Z][a-zA-Z]*)\b`, name)
		if err != nil || name == "" || !matched {
			continue
		}
		newNames = append(newNames, soundex(name))
	}
	if len(newNames) == 0 {
		fmt.Fprint(writer, pageTop, "Incorrect input", pageBottom)
		return
	}

	fmt.Fprint(writer, pageTop, formatResults(names, newNames), pageBottom)
}

//formatResults formats the result and adds it to the HTML-table.
func formatResults(names, soundexes []string) string {
	text := `<table border="1"><tr><th>Name</th><th>Soundex</th></tr>`
	for i := range names {
		text += "<tr><td>" + html.EscapeString(names[i]) + "</td><td>" +
			html.EscapeString(soundexes[i]) + "</td></tr>"
	}
	return text + "</table>"
}

//mainForm handles requests to main page.
func mainForm(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, pageTop, pageTitle, form, pageBottom)
}

func main() {
	http.HandleFunc("/", mainForm)
	http.HandleFunc("/test", viewResult)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
