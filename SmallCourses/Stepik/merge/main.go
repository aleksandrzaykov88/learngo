package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var mu sync.Mutex
	c1 := make(chan []int)
	c2 := make(chan []int)
	go func() {
		index := 0
		for i := range in1 {
			arr := make([]int, 2)
			mu.Lock()
			num1 := i
			arr[0] = index
			mu.Unlock()
			go func() {
				arr[1] = fn(num1)
				c1 <- arr
			}()
			index++
		}
		close(c1)
	}()
	go func() {
		index := 0
		for i := range in2 {
			arr := make([]int, 2)
			mu.Lock()
			num2 := i
			arr[0] = index
			mu.Unlock()
			go func() {
				arr[1] = fn(num2)
				c2 <- arr
			}()
			index++
		}
		close(c2)
	}()
	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			res1 := <-c1
			res2 := <-c2
			m[res1[0]] += res1[1]
			m[res2[0]] += res2[1]
		}
		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}
