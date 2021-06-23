package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
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

func formatInput(num float64, numArg string) string {
	var str string
	if num < 0 {
		str = strconv.FormatFloat(-num, 'f', -1, 64)
	}
	str = strconv.FormatFloat(num, 'f', -1, 64)
	if numArg == "fArg" {
		str = str + "x<sup>2</sup>"
	} else if numArg == "sArg" {
		str = str + "x"
	}
	if num > 0 && numArg != "fArg" {
		return " + " + str
	} else if num < 0 && numArg != "fArg" {
		return " - " + str
	} else {
		return ""
	}
}

func formatResult(numbers []float64, roots []complex128) string {
	a := formatInput(numbers[0], "fArg")
	b := formatInput(numbers[1], "sArg")
	c := formatInput(numbers[2], "tArg")
	return fmt.Sprintf(`%s%s%s → %v</span>`, a, b, c, roots)
}

func quadraticEquationRootsCalc(odds []float64) (complex128, complex128) {
	a := odds[0]
	b := odds[1]
	c := odds[2]
	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		ac := complex(odds[0], 0)
		bc := complex(odds[1], 0)
		Dc := complex(math.Pow(b, 2)-4*a*c, 0)
		return (-(bc) + cmplx.Sqrt(Dc)) / (2 * ac), (-(bc) - cmplx.Sqrt(Dc)) / (2 * ac)
	}
	return complex((-(b)+math.Sqrt(D))/(2*a), 0), complex((-(b)-math.Sqrt(D))/(2*a), 0)
}

func getRoots(numbers []float64) (roots []complex128) {
	x1, x2 := quadraticEquationRootsCalc(numbers)
	var nums []complex128
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
			roots := getRoots(numbers)
			fmt.Fprint(writer, formatResult(numbers, roots))
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
