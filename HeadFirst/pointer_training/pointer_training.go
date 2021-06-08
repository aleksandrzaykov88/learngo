package main

import (
	"fmt"
)

func double(number *int) {
	*number *= 2
}

func main() {
	/*var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)*/

	amount := 6
	double(&amount)
	fmt.Println(amount)
}
