package main

import (
	"fmt"
	"math"
)

//maximum returns max number from nums-slice.
func maximum(nums ...float64) float64 {
	max := math.Inf(-1)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(maximum(10, 11, 123, 1312.3))
}
