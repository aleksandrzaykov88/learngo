package main

import (
	"fmt"
	"math"
)

//maxInList is a function with a variable number of parameters that finds the largest number in the list.
func maxInList(nums ...int) int {
	max := math.Inf(-1)
	for _, num := range nums {
		if float64(num) > max {
			max = float64(num)
		}
	}
	return int(max)
}

func main() {
	fmt.Println(maxInList(37, 2, 234, -11, 11111))
}
