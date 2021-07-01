package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//averageCalc calculates average.
func averageCalc(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

//average prints average of number set imports from command line.
func average() {
	args := os.Args[1:]
	var numbers []float64
	for _, argument := range args {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Average: %0.2f\n", averageCalc(numbers...))
}

func main() {
	average()
}
