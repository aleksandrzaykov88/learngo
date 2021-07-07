package main

import (
	"fmt"
	"time"
)

func main() {
	result1 := make(chan int)
	result2 := make(chan int)

	t := time.Now()
	fmt.Printf("Start: %s\n", t.Format(time.RFC3339))

	go calculateSmthn(1000, result1)
	go calculateSmthn(2000, result2)

	fmt.Println(<-result1)
	fmt.Println(<-result2)

	fmt.Printf("Time: %s\n", time.Since(t))
}

func calculateSmthn(n int, res chan int) {
	t := time.Now()

	result := 0

	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Result: %d; Time passed: %s\n", result, time.Since(t))
	res <- result
}
