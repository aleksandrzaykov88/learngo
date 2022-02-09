package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20000; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
