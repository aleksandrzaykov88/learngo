package main

import (
	"flag"
	"fmt"
)

func main() {
	var count = flag.Int("count", 10, "test flag")
	flag.Parse()
	fmt.Println(*count)
}
