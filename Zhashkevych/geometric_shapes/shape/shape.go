package main

import (
	circle "github.com/aleksandrzaykov88/learngo/Zhashkevych/geometric_shapes/circle"
	rectangle "github.com/aleksandrzaykov88/learngo/Zhashkevych/geometric_shapes/rectangle"
	triangle "github.com/aleksandrzaykov88/learngo/Zhashkevych/geometric_shapes/triangle"
)

type GeometricShape interface {
	GetArea() float64
	GetPerimeter() float64
	SetSize()
	SetSizeFromKeyboard(string) float64
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
	var t = triangle.Triangle{}
	var r = rectangle.Rectangle{}
	CreateShape(&c)
	CreateShape(&t)
	CreateShape(&r)
}
