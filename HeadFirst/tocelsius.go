//This function converts fahrenheit temperature into celsius
package main

import (
	"fmt"
	"log"

	"github.com/aleksandrzaykov88/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degees Celsius\n", celsius)
}
