package main

import (
	"fmt"
	"io"
)

type MyReader struct{}

//Read emits ifinite 'A'-stream.
func (r MyReader) Read(b []byte) (int, error) {
	count := 0
	for i := range b {
		b[i] = 'A'
		count++
	}
	return count, nil
}

func main() {
	r := MyReader{}

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
