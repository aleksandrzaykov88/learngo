package main

import "fmt"

//feets2meters converts feets to meters.
func feets2meters(feets float64) float64 {
	return feets * 0.3048
}

func main() {
	fmt.Print("Enter the number of feets : ")
	var input float64
	fmt.Scanf("%f", &input)

	output := feets2meters(input)

	fmt.Println("In meters =", output)
}
