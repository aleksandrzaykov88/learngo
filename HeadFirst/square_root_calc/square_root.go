package main

import (
	"fmt"
	"math"
)

//squareRoot returns the square root of input number (in real, not in complex field).
func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}

func main() {
	num, _ := squareRoot(100)
	fmt.Println(num)
}
