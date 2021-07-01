package main

import (
	"fmt"
	"math"

	"github.com/aleksandrzaykov88/learngo/SimpleJavaTasks/math_functions"
)

//expTaylorE calculates the taylor expansion of exponent with a given precision.
func expTaylorE(x, e float64) float64 {
	var result float64
	result = 0
	var interValue float64
	n := 0
	for {
		interValue = math.Pow(x, float64(n)) / float64(math_functions.Factorial(n))
		result += interValue
		if math.Abs(interValue) < e {
			return result
		}
		n++
	}
}

func main() {
	fmt.Println(expTaylorE(-5, 0.0001))
}
