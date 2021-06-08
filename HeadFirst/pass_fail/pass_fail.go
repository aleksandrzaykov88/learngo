//pass_fail tells us about success or unsuccess of student's examinations
package main

import (
	"fmt"
	"log"

	"github.com/aleksandrzaykov88/packagesgo/headfirstgo/keyboard"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
