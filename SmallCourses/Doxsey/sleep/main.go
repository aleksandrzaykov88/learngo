package main

import (
	"fmt"
	"time"
)

//sleeper is a time.Sleep()-analog.
func sleeper(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

//pinger send information in c-channel.
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

//printer prints information getted from c-channel.
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		sleeper(2)
	}
}
func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
