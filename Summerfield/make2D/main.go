package main

import "fmt"

//make2D create matrix from slice with selected number of cols.
func make2D(arr []int, colsNum int) [][]int {
	k := 0
	rowsNum := len(arr) / colsNum
	if len(arr)%colsNum != 0 {
		rowsNum++
	}
	rows := make([][]int, rowsNum)
	for i := range rows {
		rows[i] = make([]int, colsNum)
		for j := range rows[i] {
			if k > len(arr)-1 {
				rows[i][j] = 0
			} else {
				rows[i][j] = arr[k]
				k++
			}
		}
	}
	return rows
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(make2D(arr, 6))
}
