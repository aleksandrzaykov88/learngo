package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Start: %s\n", t.Format(time.RFC3339))

	go func() {
		for {
			for _, r := range `—\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go calculateSmthn(1000)
	go calculateSmthn(1000)

	time.Sleep(8 * time.Second)
	fmt.Printf("Time: %s\n", time.Since(t))
}

func calculateSmthn(n int) {
	t := time.Now()

	result := 0

	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Result: %d; Time passed: %s\n", result, time.Since(t))
}
