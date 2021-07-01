package main

import (
	"fmt"
	"math"
)

//quadraticEquationRootsCalc solves the quadratic equation in real numbers.
func quadraticEquationRootsCalc(a, b, c, D float64) (float64, float64) {
	return (-(b) + math.Sqrt(D)) / (2 * a), (-(b) - math.Sqrt(D)) / (2 * a)
}

//quadraticEquationRoots gets on its input three real numbers.
//It is coefficients of quadratic equation.
//Program finds those roots and returns their amount.
func quadraticEquationRoots(a, b, c float64) {
	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		fmt.Println("There are no real roots")
	} else if D == 0 {
		x1, _ := quadraticEquationRootsCalc(a, b, c, D)
		fmt.Println("There is only one real root:", x1)
	} else {
		x1, x2 := quadraticEquationRootsCalc(a, b, c, D)
		fmt.Println("There are two real roots:", x1, x2)
	}
}

func main() {
	quadraticEquationRoots(6, 1, -2)
}
