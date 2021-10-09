package main

import (
	"fmt"
	"os"
)

//os_args prints arguments from command in several options.
func os_args() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

func main() {
	os_args()
}
