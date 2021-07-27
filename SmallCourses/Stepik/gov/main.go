package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Classificate struct {
	Id int `json:"global_id"`
}

type Classificates struct {
	Classes []Classificate
}

func main() {
	file, err := os.Open("data-20190514T0100.json")
	data, err := ioutil.ReadAll(file)
	classificates := Classificates{}
	err = json.Unmarshal(data, &classificates)
	if err != nil {
		log.Panic(err)
	}
	sum := 0
	for _, class := range classificates.Classes {
		sum += class.Id
	}
	fmt.Println(sum)
}
