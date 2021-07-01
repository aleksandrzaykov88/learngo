//Tabulation of a function sin(x) within the specified limits with a specified step

package main

import (
	"fmt"
	"math"
)

func tabSinX(x1, x2, step float64) {
	for ; x1 <= x2; x1 += step {
		fmt.Println(math.Sin(x1))
	}
}
