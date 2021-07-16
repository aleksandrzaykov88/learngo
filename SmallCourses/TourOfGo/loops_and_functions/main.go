package main

import (
	"fmt"
	"math"
)

//ErrNegativeSqrt is an error type.
type ErrNegativeSqrt float64

//Error is an error-method.
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Square root args cannot be negative in R. Your arg is: %v", float64(e))
}

//Sqrt is a custom sqrt implementation.
func Sqrt(x float64, prec float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z float64 = 1
	var step float64
	for i := 0; i < 10; i++ {
		step = (z*z - x) / (2 * z)
		z -= step
		if math.Abs(step) < prec {
			return z, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2, 1))
}
