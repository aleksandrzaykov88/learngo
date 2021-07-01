package main

import "fmt"

//negate prints reverse boolean value using the pointer mechanics.
func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
	fmt.Println(*myBoolean)
}

//double prints twice the integer value using the pointer mechanics.
func double(number *int) {
	*number *= 2
	fmt.Println(*number)
}

func main() {
	t := true
	two := 2
	negate(&t)
	double(&two)
}
