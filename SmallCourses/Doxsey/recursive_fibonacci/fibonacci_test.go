package main

import "testing"

func TestFibonacci(t *testing.T) {
	var num int
	num = fibonacci(9)
	if num != 34 {
		t.Error("Expected 34, got ", num)
	}
}
