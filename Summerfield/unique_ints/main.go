package main

import "fmt"

//UniqueInts removes doubles from input slice.
func UniqueInts(ints []int) []int {
	var support = make(map[int]int, len(ints))
	var arr = make([]int, 0)
	for i, v := range ints {
		if _, ok := support[v]; !ok {
			support[v] = i
			fmt.Println(arr)
			arr = append(arr, v)
		}
	}
	return arr
}

func main() {
	ints := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	fmt.Println(UniqueInts(ints))
}
