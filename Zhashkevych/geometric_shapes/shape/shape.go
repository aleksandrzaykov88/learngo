package shape

import (
	"github.com/aleksandrzaykov88/learngo/Zhashkevych/geometric_shapes/circle"
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
	var circle := circle.Circle{}
	CreateShape()
}
