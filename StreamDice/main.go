package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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

func diceRoller(map[int]int) {
	/*
		var roll map[int][]int
		rand.Seed(time.Now().UnixNano())
		for element := range map {
			for i:=1; i< value; i++ {
			var dice = NewDice(key)
			roll[key] = key
			roll[value] = append(roll[value], dice.Roll())
			}
		}*/
}

func diceHandler(rolls string) {
	var roll map[int]int
	//parse string by ;
	//parse strings by :
	//return map int/int
}

func StreamDice(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("sendedData")
	fmt.Println("Receive ajax post data string ", data)

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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe("localhost:8001", nil)
}
