//This function prints Fibonacci row till n-member
package main

import (
	"fmt"
)

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
