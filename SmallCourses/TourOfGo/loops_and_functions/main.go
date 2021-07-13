package main

import (
	"fmt"
	"math"
)

//Sqrt is a custom sqrt implementation.
func Sqrt(x float64, prec float64) float64 {
	var z float64 = 1
	var step float64
	for i := 0; i < 10; i++ {
		step = (z*z - x) / (2 * z)
		z -= step
		if math.Abs(step) < prec {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 0.00002) - math.Sqrt(2))
}
