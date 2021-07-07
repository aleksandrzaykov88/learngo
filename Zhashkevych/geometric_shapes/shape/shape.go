package main

import (
	circle "github.com/aleksandrzaykov88/learngo/Zhashkevych/geometric_shapes/circle"
)

type GeometricShape interface {
	GetArea() float64
	GetPerimeter() float64
	SetSize()
	SetSizeFromKeyboard() float64
	Show()
	Check() error
}

//CreateShape tests geometric shape interface.
func CreateShape(g GeometricShape) {
	g.SetSize()
	g.Show()
}

func main() {
	var c = circle.Circle{}
	CreateShape(&c)
}
