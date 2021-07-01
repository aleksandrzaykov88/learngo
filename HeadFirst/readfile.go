package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//readFile reads .txt file and print it line by line.
func readFile() {
	file, err := os.Open("C:/Users/admin/Documents/azaykov/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
