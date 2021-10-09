package main

import "fmt"

//Flatten convers two-dimensional array in one-dimension array.
func Flatten(matrix [][]int) []int {
	var result = make([]int, 0)
	for i := range matrix {
		for _, val := range matrix[i] {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	irregularMatrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	slice := Flatten(irregularMatrix)
	fmt.Printf("1x%d: %v\n", len(slice), slice)
}
