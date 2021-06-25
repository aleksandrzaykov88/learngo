package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
		log.Fatal(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
				log.Fatal(err)
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("C:/Users/admin/Documents/azaykov/learngo")
	if err != nil {
		log.Fatal(err)
	}
}
