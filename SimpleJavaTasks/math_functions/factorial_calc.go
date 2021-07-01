//Package math_functions provides users possibility to use some popular math functions.
package math_functions

//factorial returns factorial of input n
func Factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
