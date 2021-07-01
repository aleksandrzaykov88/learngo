package main

import (
	"errors"
	"fmt"
	"log"
)

//divide returns the result of dividing the first argument by the second.
func divide(divident float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return divident / divisor, nil
}

func main() {
	result, err := divide(100, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
