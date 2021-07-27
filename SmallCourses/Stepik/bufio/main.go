package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	buf := make([]byte, 10)
	n, err := rd.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("прочитано %d байт: %s\n", n, buf)
	var pos int
	for i := 0; true; i++ {
		s, _ := rd.ReadString(';')
		if s == "0;" {
			pos = i + 1
			break
		}
	}
	fmt.Println(pos)
}
