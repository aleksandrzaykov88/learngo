package main

import (
	"fmt"
	"strings"
)

//WordCount counts words in sentence.
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	return wordMap
}

func main() {
	fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
}
