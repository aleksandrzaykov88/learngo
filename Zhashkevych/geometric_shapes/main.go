package main

import (
	"circle"
	"fmt"
)

func main() {
	var c circle.Circle
	c.SetSize()
	fmt.Println(c.GetArea())
}
