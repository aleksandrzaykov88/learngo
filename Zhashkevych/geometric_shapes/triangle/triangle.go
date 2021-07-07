package triangle

import (
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/aleksandrzaykov88/learngo/HeadFirst/keyboard"
)

//Triangle-struct describes geometric shape triangle.
type Triangle struct {
	A float64
	B float64
	C float64
}

//Check uses triangle inequality to check user input and determine whether it is a triangle or not.
func (t *Triangle) Check() error {
	if t.A < 0 || t.B < 0 || t.C < 0 {
		return errors.New("Side size can't be less than zero!")
	}
	if t.A > t.B+t.C || t.B > t.C+t.A || t.C > t.A+t.B {
		return errors.New("This is not a triangle!")
	}
	return nil
}

//SetSidesFromKeyboard sets sizes of triangle sides from user input.
func (t *Triangle) SetSizeFromKeyboard(sideName string) float64 {
	fmt.Printf("Enter %s side: ", sideName)
	number, err := keyboard.GetFloat()
	if err != nil {
		return 0
	}
	return number
}

//SetSides sets sizes of triangle sides.
func (t *Triangle) SetSize() {
	t.A = t.SetSizeFromKeyboard("A")
	t.B = t.SetSizeFromKeyboard("B")
	t.C = t.SetSizeFromKeyboard("C")
	err := t.Check()
	if err != nil {
		log.Fatal(err)
		return
	}
}

//GetArea returns the area of triangle using the Heron's formula.
func (t *Triangle) GetArea() float64 {
	sP := t.GetPerimeter() / 2
	fmt.Println(sP)
	return math.Sqrt(sP * (sP - t.A) * (sP - t.B) * (sP - t.C))
}

//GetPerimeter returns the perimeter of triangle.
func (t *Triangle) GetPerimeter() float64 {
	return t.A + t.B + t.C
}

//Show prints result of calculating area and perimeter of triangle.
func (t *Triangle) Show() {
	tArea := t.GetArea()
	tPerimeter := t.GetPerimeter()
	fmt.Printf("Rectangle area: %f cm^2", tArea)
	fmt.Println()
	fmt.Printf("Circumference of the rectangle: %f cm", tPerimeter)
}
