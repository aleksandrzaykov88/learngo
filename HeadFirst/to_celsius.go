package main

import (
	"fmt"
	"log"
)

//toCelsius() converts fahrenheit temperature into celsius.
func toCelsius() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degees Celsius\n", celsius)
}
