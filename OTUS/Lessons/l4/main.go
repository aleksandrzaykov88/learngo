package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5}
	c := []int{6, 7}

	fmt.Println(Concat(a, b, c))
}

func Concat(slices ...[]int) []int {
	newSlice := make([]int, 0)
	for _, slice := range slices {
		newSlice = append(newSlice, slice...)
	}

	return newSlice
}
