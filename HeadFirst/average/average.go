package main

import (
	"fmt"
	"log"

	"github.com/aleksandrzaykov88/learngo/HeadFirst/datafile"
)

//average prints average of number set imports from file.
func average() {
	numbers, err := datafile.GetFloats("C:/Users/admin/Documents/azaykov/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

func main() {
	average()
}
