package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	go parseURL("http://example.com/")
	parseURL("http://youtube.com/")

	fmt.Printf("Parsing completed. Time Elapesed: %.2f seconds", time.Since(t).Seconds())
}

func parseURL(url string) {
	for i := 0; i < 5; i++ {
		latency := rand.Intn(500) + 500

		time.Sleep(time.Duration(latency) * time.Millisecond)

		fmt.Printf("Parsing <%s> - Step %d - Latency %d\n", url, i+1, latency)
	}
}
