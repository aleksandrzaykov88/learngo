package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	rand.Seed(time.Now().UTC().UnixNano())

	ret := calculator(arguments, done)
	for i := 1; i <= 10; i++ {
		arguments <- rand.Intn(99)
	}
	close(done)
	fmt.Println(<-ret)
}
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {

	c := make(chan int)

	go func() {
		sum := 0
		for {
			select {
			case num := <-arguments:
				sum += num
			case <-done:
				c <- sum
				defer close(c)
				return
			}
		}
	}()

	return c
}
