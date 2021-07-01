package main

import "fmt"

//linearSystemSolutionsNumberCheck checks the number of solutions of a system of linear equations
func linearSystemSolutionsNumberCheck(a1, b1, c1, a2, b2, c2 float64) string {
	if a1*b2 != a2*b1 {
		fmt.Println("There is only one solution of this system.")
		return "One"
	} else if (a1*b2 == a2*b1) && ((a1*c2 != a2*c1) || (b1*c2 != b2*c1)) {
		fmt.Println("There are no solutions of this system.")
		return "No"
	} else if (a1*b2 == a2*b1) && (a1*c2 == a2*c1) && (b1*c2 == b2*c1) {
		fmt.Println("The system has infinitely many solutions.")
		return "Many"
	}
	return "Err"
}

//linearSystemEquationCalc solves a system of two linear equations in real numbers using the Gauss Elimination.
func linearSystemEquationCalc(matrix [2][2]float64, freeMemb [2]float64) {
	if amount := linearSystemSolutionsNumberCheck(matrix[0][0], matrix[0][1], freeMemb[0], matrix[1][0], matrix[1][1], freeMemb[1]); amount == "No" {
		return
	}
	n := len(freeMemb)
	x := make([]float64, n)
	for k := 0; k <= (n - 2); k++ {
		for i := (k + 1); i < n; i++ {
			if matrix[i][k] == 0 {
				continue
			}
			factor := matrix[k][k] / matrix[i][k]
			for j := k; j < n; j++ {
				matrix[i][j] = matrix[k][j] - matrix[i][j]*factor
			}
			freeMemb[i] = freeMemb[k] - freeMemb[i]*factor
		}
	}
	x[n-1] = freeMemb[n-1] / matrix[n-1][n-1]
	for i := (n - 2); i >= 0; i-- {
		var sum_ax float64 = 0
		for j := (i + 1); j < n; j++ {
			sum_ax += matrix[i][j] * x[j]
		}
		x[i] = (freeMemb[i] - sum_ax) / matrix[i][i]
	}
	fmt.Println(x)
}

func main() {
	linearSystemEquationCalc([2][2]float64{{2, -3}, {6, -9}}, [2]float64{7, 12})
}
