package main

import (
	"fmt"
	"log"
)

//passOrFail() tells us about success or unsuccess of student's examinations.
func passOrFail() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
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
