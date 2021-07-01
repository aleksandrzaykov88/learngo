package main

import "fmt"

func recursionCounter(start, end int) {
	fmt.Println(start)
	if start < end {
		recursionCounter(start+1, end)
	}
}
