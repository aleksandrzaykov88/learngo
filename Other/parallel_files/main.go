package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lines := make(chan string, 1000)
	results := make(chan string, workers*4)
	done := make(chan struct{}, workers)

	files := []string{"test.txt", "test1.txt"}
	for _, f := range files {
		go func(file string) {
			reafFile(file, lines)
			processLines(done, lines, results)
			waitUntil(done)
		}(f)
	}
	for {
		fmt.Println(<-results)
	}
}

func reafFile(filename string, lines chan string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	strings := strings.Split(string(input), "\n")
	for _, str := range strings {
		lines <- str
	}
}

func processLines(done chan<- struct{}, lines <-chan string, results chan string) {
	for i := 0; i < workers; i++ {
		go func() {
			for line := range lines {
				results <- line
			}
			done <- struct{}{}
		}()
	}
}

func waitUntil(done <-chan struct{}) {
	for i := 0; i < workers; i++ {
		<-done
	}
}
