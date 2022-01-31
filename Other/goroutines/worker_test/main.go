package main

import (
	"fmt"
	"math/rand"
	"time"
)

func counter() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var i int64
	for i = 0; i < r1.Int63n(100000000000000000); i++ {
		i++
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 0; r <= 6; r++ {
		fmt.Println("Result received from worker: ", <-results)
	}
}

func worker(ID int, jobs chan int, results chan int) {
	for job := range jobs {
		fmt.Println("Worker ", ID, " is working on job ", job)
		start := time.Now()
		counter()
		duration := time.Since(start)
		fmt.Println("Worker ", ID, " completed work on job ", job, " within ", duration)
	}
}
