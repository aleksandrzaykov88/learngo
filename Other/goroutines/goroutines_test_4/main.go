package main

import (
	"fmt"
	"time"
)

func main() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			message1 <- "Прошло пол секунды"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			message2 <- "Прошло 2 секунды"
		}
	}()

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		}
	}
}
