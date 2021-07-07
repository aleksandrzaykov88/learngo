package main

import (
	"fmt"
	"github.com/aleksandrzaykov88/learngo/HeadFirst/keyboard"
)

type GeometricShape interface {
	GetArea() float64
	GetPerimeter() float64
	SetSize()
	SetSizeFromKeyboard()
	Show()
	Check() error
}

//CreateShape tests geometric shape interface.
func CreateShape(g GeometricShape) {
	g.SetSize()
	g.Show()
}

func main() {
	f, _ := keyboard.GetFloat()
	fmt.Println(f)
}
