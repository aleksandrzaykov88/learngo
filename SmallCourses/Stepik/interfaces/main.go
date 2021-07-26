package main

import (
	"fmt"
)

func main() {
	var value1 interface{} = 4.5
	var value2 interface{} = 3.0
	var operation interface{} = "-"
	switch operation {
	case "+":
		fmt.Printf("%.4f", value1.(float64)+value2.(float64))
	case "-":
		fmt.Printf("%.4f", value1.(float64)-value2.(float64))
	case "/":
		fmt.Printf("%.4f", value1.(float64)/value2.(float64))
	case "*":
		fmt.Printf("%.4f", value1.(float64)*value2.(float64))
	default:
		fmt.Println("неизвестная операция")
		return
	}
}

func checkFloat(n interface{}) bool {
	switch v := n.(type) {
	case float64:
		return true
	default:
		fmt.Printf("value=%v: %T!\n", v, v)
		return false
	}
}

func checkString(n interface{}) bool {
	switch n.(type) {
	case string:
		return true
	default:
		fmt.Println("неизвестная операция")
		return false
	}
}
