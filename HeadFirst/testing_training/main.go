package main

import (
	"fmt"
	"strings"
)

//joinWithCommas add words separating by commas in the end of the phrase.
//if amount of additional phrases < 3 there is some special rules.
func joinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", joinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", joinWithCommas(phrases))
}
