package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type JSONResults struct {
	Throwed string `json:"throwed"`
	Sum     int    `json:"sum"`
	D100    string `json:"d100"`
	D20     string `json:"d20"`
}

func sum() {
	//dont sum d100 and d20
}

func min() {

}

func max() {

}

func average() {

}

func createJSON() {

}

func diceRoller(rolls map[int]int) map[string]int {
	rand.Seed(time.Now().UnixNano())
	var sum int
	var result = make(map[string]int)
	for k, v := range rolls {
		sum = 0
		var dice = NewDice(k)
		for i := 1; i <= v; i++ {
			sum += dice.Roll()
		}
		newKey := "d" + fmt.Sprint(k)
		result[newKey] = sum
	}
	return result
}

func diceHandler(rolls string) map[int]int {
	var rollResults = make(map[int]int)
	s := strings.Split(rolls, ";")
	for _, line := range s {
		if strings.HasPrefix(line, "d") {
			sLine := strings.Split(line, ":")
			k, err := strconv.Atoi(sLine[0][1:])
			if err != nil {
				log.Fatal(err)
			}
			v, err := strconv.Atoi(sLine[1])
			if err != nil {
				log.Fatal(err)
			}
			rollResults[k] = v
		}
	}
	return rollResults
}

func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("sendedData")
	var rollResults = make(map[string]int)
	if data != "" {
		rollMap := diceHandler(data)
		rollResults = diceRoller(rollMap)
		fmt.Println(rollResults)

		test := &JSONResults{Throwed: "test", Sum: 1, D100: "dfsf", D20: "tt"}

		b, err := json.Marshal(test)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(b))
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func StreamDice(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", "")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/img/favicon.ico")
}

func main() {
	http.HandleFunc("/", StreamDice)
	http.HandleFunc("/ajax", AjaxHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe("localhost:8001", nil)
}
