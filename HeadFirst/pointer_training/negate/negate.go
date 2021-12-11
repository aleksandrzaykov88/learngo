package negate

import "fmt"

//negate() prints reverse boolean using the pointer mechanics.
func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
	fmt.Println(*myBoolean)
}
