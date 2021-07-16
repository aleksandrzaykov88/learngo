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
	Log     string `json:"log"`
	Name    string `json:"throwman"`
}

//GetThrowman gets player-name from ajax.
func GetThrowman(rolls string) string {
	s := strings.Split(rolls, ";")
	result := ""
	for _, line := range s {
		if strings.HasPrefix(line, "name") {
			sLine := strings.Split(line, ":")
			result = sLine[1]
		}
	}
	return result
}

//StringRolls generates result-string for d20&d100 dices, because they should not be counted in sum.
func StringRolls(rollMap map[int]int) string {
	rand.Seed(time.Now().UnixNano())
	result := ""
	for k, v := range rollMap {
		if k != 20 && k != 100 {
			continue
		}
		var dice = NewDice(k)
		for i := 1; i <= v; i++ {
			result += fmt.Sprint(dice.Roll()) + " "
		}
	}
	return result
}

//Sum returns throw-result.
func Sum(rollResults map[string]int) int {
	result := 0
	for k, v := range rollResults {
		if k == "d20" || k == "d100" {
			continue
		}
		result += v
	}
	return result
}

//Logger returns information message for humans.
func Logger(rollMap map[int]int) string {
	result := ""
	addString := ""
	for k, v := range rollMap {
		addString = "The D" + fmt.Sprint(k) + " die has been rolled " + fmt.Sprint(v) + " time"
		result += addString
		if v != 1 {
			result += "s"
		}
		result += "!<br>"
	}
	return result
}

//Throwed returns information message for json.
func Throwed(rollMap map[int]int) string {
	result := ""
	addString := ""
	for k, v := range rollMap {
		addString = "D" + fmt.Sprint(k) + "×" + fmt.Sprint(v) + ";\n"
		result += addString
	}
	return result
}

//createJSON generate json-responce.
func createJSON(rolls string) *JSONResults {
	rollMap := diceHandler(rolls)
	rollResults := diceRoller(rollMap)

	jsonResult := &JSONResults{}
	jsonResult.Throwed = Throwed(rollMap)
	jsonResult.Sum = Sum(rollResults)
	jsonResult.D100 = StringRolls(rollMap)
	jsonResult.D20 = StringRolls(rollMap)
	jsonResult.Log = Logger(rollMap)
	jsonResult.Name = GetThrowman(rolls)

	return jsonResult
}

//diceRoller roll dices and returns roll-results.
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

//diceHandler parses input string from ajax.
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

//AjaxHandler handles ajax-query.
func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("sendedData")
	if data != "" {
		jsonResult := createJSON(data)

		b, err := json.Marshal(jsonResult)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(b))
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

//StreamDice handles main-page of app.
func StreamDice(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", name)
}

//faviconHandler shows favicon.
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
