package frequencyanalysis

import (
	"sort"
	"strings"
)

func TopTenWords(text string) []string {
	var counter = make(map[string]int)
	var res []string
	var frequencies []int
	var excludedValues []int

	words := strings.Fields(text)

	for _, word := range words {
		counter[word]++
	}

	for _, v := range counter {
		if !contains(excludedValues, v) {
			frequencies = append(frequencies, v)
		}
		excludedValues = append(excludedValues, v)
	}

	sort.Ints(frequencies)

	for i := len(frequencies) - 1; i >= 0; i-- {
		supportSlice := make([]string, 0)
		for k, v := range counter {
			if v == frequencies[i] {
				supportSlice = append(supportSlice, k)
			}
		}
		sort.Strings(supportSlice)
		res = append(res, supportSlice...)
	}

	return res
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
