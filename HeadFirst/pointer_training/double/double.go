package double

import "fmt"

//double prints twice the value using pointer mechanics.
func double(number *int) {
	*number *= 2
	fmt.Println(*number)
}
