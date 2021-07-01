package main

import (
	"fmt"
	"math"
)

//tabSinX tabulates the sin(x) function within the specified limits with a specified step.
func tabSinX(x1, x2, step float64) {
	for ; x1 <= x2; x1 += step {
		fmt.Println(math.Sin(x1))
	}
}

func main() {
	tabSinX(0, 5, 0.2)
}
