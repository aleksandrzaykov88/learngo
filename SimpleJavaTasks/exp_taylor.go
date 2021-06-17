package main

import (
	"math"
)

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//expTaylorE() calculates the taylor expansion of exponent with a given precision.
func expTaylorE(x, e float64) float64 {
	var result float64
	result = 0
	var interValue float64
	n := 0
	for {
		interValue = math.Pow(x, float64(n)) / float64(factorial(n))
		result += interValue
		if math.Abs(interValue) < e {
			return result
		}
		n++
	}
}
