package main

import (
	"fmt"
	"math"
)

//powTwo prints the powers of 2 from 1 to n.
func powTwo(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(math.Pow(2, float64(i)))
	}
}

func main() {
	powTwo(18)
}
