package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var b1 byte = 0b11000010
	var b2 byte = 0b10100011
	var buf []byte
	buf = append(buf, b1)
	buf = append(buf, b2)
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("%c\n", r)
}
