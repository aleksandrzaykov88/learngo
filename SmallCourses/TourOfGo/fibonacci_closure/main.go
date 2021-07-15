package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var result int
	var z int
	x := 0
	y := 1
	return func() int {
		z = x + y
		result = z - y
		x = y
		y = z
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
