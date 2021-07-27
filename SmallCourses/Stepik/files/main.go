package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//walkFunc implements recurvie walk in file and finds some number.
func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}

	if strings.HasSuffix(info.Name(), "zip") {
		r, err := zip.OpenReader(info.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		for _, f := range r.File {
			file, _ := f.Open()
			r := csv.NewReader(file)
			record, _ := r.ReadAll()
			if len(record) == 10 {
				fmt.Println(record[4][2])
			}
		}
	}

	return nil
}

func main() {
	const root = "."
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
