package main

import (
	"errors"
	"fmt"
	"math"
)

const pi = math.Pi

func main() {
	printCircleArea(2)
}

//printCircleArea prints result of calculating area of circle.
func printCircleArea(radius int) {
	circleArea, err := calculateCirceArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Circle radius: %d cm", radius)
	fmt.Println()
	fmt.Printf("Circle area: %f cm^2", circleArea)
}

//calculateCirceArea calculates the area of a circle with a given radius.
func calculateCirceArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Radius cannot be negative!")
	}
	return float32(radius) * float32(radius) * pi, nil
}
