package main

import (
	"fmt"
	"sort"
)

//inRange returns sorted slice of args between min and max.
func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func main() {
	fmt.Println(inRange(1, 10, 34, 4, 52, 4, 5, 2, 1.3))
}
