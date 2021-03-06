package main

import (
	"fmt"
	"sort"
)

//threeNumsCalc takes on input three values and return their product, average and sorts them by ascension.
func threeNumsCalc(a, b, c int) {
	af := float64(a)
	bf := float64(b)
	cf := float64(c)
	product := af * bf * cf
	sum := af + bf + cf
	sortedNums := []float64{af, bf, cf}
	sort.Float64s(sortedNums)
	fmt.Println("Product of numbers:", a, "*", b, "*", c, "=", product)
	fmt.Println("Average of numbers:", a, b, c, "is", sum/3)
	fmt.Println("Sorting numbers:", a, b, c, "by ascension:", sortedNums)
}

func main() {
	threeNumsCalc(1, 2, 5)
}
