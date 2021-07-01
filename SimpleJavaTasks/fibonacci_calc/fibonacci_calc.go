package main

import (
	"fmt"
)

//fibonacci prints the Fibonacci row till n-member.
func fibonacci(n int) {
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		var z = x + y
		fmt.Println(z - y)
		x = y
		y = z
	}
}

func main() {
	fibonacci(10)
}
