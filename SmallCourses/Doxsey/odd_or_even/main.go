package main

import "fmt"

//oddOrEven takes a number, divides it in half and returns true if the original number is even, and false if it is odd
func oddOrEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(oddOrEven(37))
	fmt.Println(oddOrEven(-2))
}
