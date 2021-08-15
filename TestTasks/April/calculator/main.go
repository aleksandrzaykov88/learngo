package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//IsBalanced returns true if brackets are balanced.
func IsBalanced(text string) bool {
	isBalanced := true
	s := make([]rune, 0, len(text))

	for _, c := range text {
		if c == '(' {
			s = append(s, c)
		} else if c == ')' {
			if len(s) == 0 {
				isBalanced = false
				break
			}
			s = s[:len(s)-1]
		}
	}

	if len(s) != 0 {
		isBalanced = false
	}

	return isBalanced
}

//CalcHandle handles input URL.
func CalcHandle(w http.ResponseWriter, r *http.Request) {
	queue := make([]float64, 0)

	//Regexp for parse all expression.
	re := regexp.MustCompile(`[+-]?\d+(\.\d+)?|[[+|-]?[(|)]|[+|-|*|\/?]|[+|-]`)
	//Regexp for find errors in expression.
	invalid := regexp.MustCompile(`^[*|\/]|[+|\-|*|\/]$|[+|\-|*|\/]{3,}`)
	//Regexp for find nums
	numberExp := regexp.MustCompile(`^[+-]?\d+(\.\d+)?$`)
	expression, ok := r.URL.Query()["expr"]
	if !ok || len(expression[0]) < 1 {
		return
	}
	expr := strings.Replace(expression[0], " ", "+", 1)

	//If brackets are balanced or
	//*|/ in start of expression or
	//Some operator at the end of expression or
	//3+ operators in a row,
	//Then it is a invalid expression.
	if !IsBalanced(expr) || len(invalid.FindAllString(expr, -1)) > 0 {
		fmt.Println("Invalid expression.")
		return
	}

	for _, token := range re.FindAllString(expr, -1) {
		//If token is a number - i add it to queue.
		if numberExp.MatchString(token) {
			num, _ := strconv.ParseFloat(token, 64)
			queue = append(queue, num)
		}
	}

	fmt.Println(queue)
}

func main() {
	http.HandleFunc("/", CalcHandle)
	http.ListenAndServe("localhost:8080", nil)
}
