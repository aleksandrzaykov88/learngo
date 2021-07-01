package main

import (
	"fmt"
	"log"

	"github.com/aleksandrzaykov88/learngo/HeadFirst/datafile"
)

//countStrings calculates amount of occurrences of specific string in file.
func countStrings() {
	lines, err := datafile.GetStrings("C:/Users/admin/Documents/azaykov/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}

func main() {
	countStrings()
}
