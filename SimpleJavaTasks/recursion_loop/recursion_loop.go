package main

import "fmt"

//recursionCounter is an implementation of loop by recursion.
func recursionCounter(start, end int) {
	fmt.Println(start)
	if start < end {
		recursionCounter(start+1, end)
	}
}

func main() {
	recursionCounter(1, 10)
}
