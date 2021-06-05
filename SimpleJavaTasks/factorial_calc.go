//This function returns factorial of input n
package main

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
