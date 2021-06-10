package main

import "fmt"

//linearSystemSolutionsNumberCheck ()checks the number of solutions of a system of linear equations
func linearSystemSolutionsNumberCheck(a1, b1, c1, a2, b2, c2 float64) string {
	if a1*b2 != a2*b1 {
		fmt.Println("There is only one solution of this system")
		return "One"
	} else if (a1*b2 == a2*b1) && ((a1*c2 != a2*c1) || (b1*c2 != b2*c1)) {
		fmt.Println("There are no solutions of this system")
		return "No"
	} else if (a1*b2 == a2*b1) && (a1*c2 == a2*c1) && (b1*c2 == b2*c1) {
		fmt.Println("The system has infinitely many solutions")
		return "Many"
	}
	return "Err"
}

//linearSystemEquationCalc() solves a system of two linear equations in real numbers.
func linearSystemEquationCalc(f [3]float64, s [3]float64) {
	fmt.Println(f)
	fmt.Println(s)
	linearSystemSolutionsNumberCheck(f[0], f[1], f[2], s[0], s[1], s[2])
}
