package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			message <- fmt.Sprintf("%d", i)
			time.Sleep(500 * time.Millisecond)
		}

		close(message)
	}()

	for msg := range message {
		fmt.Println(msg)
	}
}
