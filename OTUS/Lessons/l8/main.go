package main

import (
	"fmt"
	"time"
)

func main() {
	var start = make(chan chan struct{})

	for i := 0; i < 10000; i++ {
		go func(i int) {
			<-start

			if i%1000 == 0 {
				fmt.Println(i)
			}

		}(i)
	}
	close(start)
	time.Sleep(time.Second)

	tickTock()
}

func tickTock() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticktock")
		}
	}
}
