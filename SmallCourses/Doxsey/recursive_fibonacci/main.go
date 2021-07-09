package main

import "fmt"

//fibonacci returns the Fibonacci row n-member.
func fibonacci(n int) int {
	if n == 0 || n < 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(9))
}
