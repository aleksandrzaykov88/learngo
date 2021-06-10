package main

import (
	"fmt"
	"math"
)

//squareRoot() returns square root of input number (in real, not in complex field).
func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}
