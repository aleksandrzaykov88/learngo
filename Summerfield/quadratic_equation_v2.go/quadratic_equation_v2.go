package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	pageTop = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Quadratic Equation Solver</title>
	</head>
	<body>
		<h1>Quadratic Equation Solver</h1>
		<p>Solves equations of the form ax<sup>2</sup> + bx + c</p>`
	form = `<form class="calc_form" action="/" method="POST">
	<input class="num_field" type="text" name="coef1" id="coef1">
	<label for="coef1">x<sup>2</sup></label>
	<span> + </span>
	<input class="num_field" type="text" name="coef2" id="coef2">
	<label for="coef2">x</label>
	<span> + </span>
	<input class="num_field" type="text" name="coef3" id="coef3">
	<span> → </span>
	<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body>
	</html>	
	<style>
		.calc_form {
			display: flex;
			align-items: baseline;
			justify-content: space-between;
			width: 268px;
			margin-bottom: 20px;
		}
		.num_field {
			width: 20px;
			text-align: right;
		}
	</style>`
	anError = `<p class="error">%s</p>`
)

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func formatStats(stats []float64) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Roots: </td><td>%v</td></tr>
</table>`, stats)
}

func quadraticEquationRootsCalc(odds []float64) (float64, float64) {
	a := odds[0]
	b := odds[1]
	c := odds[2]
	D := math.Pow(b, 2) - 4*a*c
	return (-(b) + math.Sqrt(D)) / (2 * a), (-(b) - math.Sqrt(D)) / (2 * a)
}

func getStats(numbers []float64) (stats []float64) {
	x1, x2 := quadraticEquationRootsCalc(numbers)
	var nums []float64
	nums = append(nums, x1, x2)
	return nums
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var nums []float64
	for i := 1; i <= 3; i++ {
		key := "coef" + strconv.Itoa(i)
		if num, found := request.Form[key]; found && len(num) > 0 {
			for _, field := range strings.Fields(num[0]) {
				if x, err := strconv.ParseFloat(field, 64); err != nil {
					return nums, "'" + field + "' is invalid", false
				} else {
					nums = append(nums, x)
				}
			}
		}
	}
	if len(nums) < 3 {
		return nums, "", false
	}
	return nums, "", true
}
