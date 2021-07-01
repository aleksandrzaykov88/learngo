package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aleksandrzaykov88/learngo/HeadFirst/datafile"
)

func main() {
	numbers, err := datafile.GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)
}
