package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//oper is a type for describes the math operators and their priority.
type oper struct {
	sign     string
	priority int
}

//Hardcoded block of math operators.
var plus = oper{sign: "+", priority: 1}
var minus = oper{sign: "-", priority: 1}
var multiply = oper{sign: "*", priority: 2}
var divide = oper{sign: "/", priority: 2}

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

//IdentifyToken returns one of the possible options of input token.
func IdentifyToken(token string) oper {
	switch token {
	case plus.sign:
		return plus
	case minus.sign:
		return minus
	case multiply.sign:
		return multiply
	case divide.sign:
		return divide
	default:
		return oper{sign: token, priority: 3}
	}
}

//Evaluate applies the operator from strStack to the last two numbers of the numStack.
func Evaluate(numStack *Stack, strStack *StrStack) {
	op, _ := strStack.Pop()
	switch op {
	case "+":
		num2, _ := numStack.Pop()
		num1, _ := numStack.Pop()
		numStack.Push(num1 + num2)
	case "-":
		num2, _ := numStack.Pop()
		num1, _ := numStack.Pop()
		numStack.Push(num1 - num2)
	case "*":
		num2, _ := numStack.Pop()
		num1, _ := numStack.Pop()
		numStack.Push(num1 * num2)
	case "/":
		num2, _ := numStack.Pop()
		num1, _ := numStack.Pop()
		numStack.Push(num1 / num2)
	}
}

//CalculateExpr calculates value of input expression.
//Implementation of shunting yard algorithm.
func CalculateExpr(expr []string) float64 {
	numStack := new(Stack)
	strStack := new(StrStack)

	//Regexp for find nums.
	numberExp := regexp.MustCompile(`^[+-]?\d+(\.\d+)?$`)

	for _, token := range expr {
		if numberExp.MatchString(token) { //if token is number
			num, _ := strconv.ParseFloat(token, 64)
			numStack.Push(num)
		} else { //if token is on of +,-,/,*,(,)
			if len(*strStack) == 0 {
				strStack.Push(token)
			} else {
				t := *strStack
				stackToken := IdentifyToken(t[len(*strStack)-1])
				inputToken := IdentifyToken(token)
				//Round brackets logic.
				if inputToken.sign == ")" {
					for {
						stackToken = IdentifyToken(t[len(*strStack)-1])
						if stackToken.sign != "(" {
							Evaluate(numStack, strStack)
						} else {
							strStack.Pop()
							break
						}
					}
					continue
				}

				if inputToken.priority > stackToken.priority || stackToken.sign == "(" || stackToken.sign == ")" {
					strStack.Push(inputToken.sign)
					continue
				}
				for {

					if len(*strStack) == 0 || stackToken.sign == "(" || stackToken.sign == ")" {
						strStack.Push(inputToken.sign)
						break
					} else {
						stackToken = IdentifyToken(t[len(*strStack)-1])
					}

					stackToken = IdentifyToken(t[len(*strStack)-1])

					if inputToken.priority <= stackToken.priority && stackToken.sign != "(" && stackToken.sign != ")" {
						Evaluate(numStack, strStack)
					} else {
						strStack.Push(inputToken.sign)
						break
					}
				}
			}
		}
	}
	for {
		if len(*strStack) == 0 {
			break
		}
		Evaluate(numStack, strStack)
	}
	num, _ := numStack.Pop()
	return num
}

//CalcHandle handles input URL.
func CalcHandle(w http.ResponseWriter, r *http.Request) {
	//Regexp for parse all expression.
	re := regexp.MustCompile(`[+-]?\d+(\.\d+)?|[(|)]|[+|-|*|\/?]|[+|-]`)
	//Regexp for find errors in expression.
	invalid := regexp.MustCompile(`^[*|\/]|[+|\-|*|\/]$|[+|\-|*|\/]{3,}`)

	expression, ok := r.URL.Query()["expr"]
	if !ok || len(expression[0]) < 1 {
		return
	}
	expr := strings.Replace(expression[0], " ", "+", -1)

	//Trick for accounting unary operators.
	signs := regexp.MustCompile(`[\d|\)][\-|+]\d`)
	for {
		changeSigns := signs.FindAllString(expr, -1)
		if len(changeSigns) > 0 {
			for _, s := range changeSigns {
				expr = strings.Replace(expr, s, s[:len(s)-1]+"+"+s[len(s)-1:], -1)
			}
		} else {
			break
		}
	}

	//If brackets are balanced or
	//*|/ in start of expression or
	//Some operator at the end of expression or
	//3+ operators in a row,
	//Then it is a invalid expression.
	if !IsBalanced(expr) || len(invalid.FindAllString(expr, -1)) > 0 {
		err := fmt.Errorf("Invalid expression.")
		fmt.Println(err)
		return
	}
	fmt.Println(CalculateExpr(re.FindAllString(expr, -1)))
}

func main() {
	http.HandleFunc("/", CalcHandle)
	http.ListenAndServe("localhost:8080", nil)
}
