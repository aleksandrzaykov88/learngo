package main

import "golang.org/x/tour/pic"

//Pic displays picture, interpreting the integers as grayscale.
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for j := 0; j < dy; j++ {
		result[j] = make([]uint8, dx)
		for i := 0; i < dx; i++ {
			result[j][i] = uint8((i + j) / 2)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
