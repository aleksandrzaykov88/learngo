package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Average struct {
	Average float64
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

func main() {
	file, err := os.Open("text.json")
	data, err := ioutil.ReadAll(file)
	g := Group{}
	err = json.Unmarshal(data, &g)
	if err != nil {
		log.Panic(err)
	}
	countRank := 0
	countStudent := 0
	for _, student := range g.Students {
		countStudent++
		for rank := range student.Rating {
			if rank != 0 {
				countRank += rank / rank
			} else {
				countRank++
			}
		}
	}
	a := Average{}
	a.Average = float64(countRank) / float64(countStudent)
	res, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", res)
}
