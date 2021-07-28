package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Classificate struct {
	Globid int `json:"global_id"`
}

func main() {
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		log.Panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}
	var decoded []Classificate
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		log.Panic(err)
	}
	sum := 0
	for _, v := range decoded {
		sum += v.Globid
	}
	fmt.Println(sum)
}
