package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aleksandrzaykov88/learngo/HeadFirst/keyboard"
)

//Rectangle-struct describes geometric shape rectangle.
type Rectangle struct {
	A float64
	B float64
}

//IsRectangle check user input for negative values of side.
func (r *Rectangle) IsRectangle() error {
	if r.A < 0 || r.B < 0 {
		return errors.New("Side size can't be less than zero!")
	}
	return nil
}

//SetSidesFromKeyboard sets sizes of rectangle sides from user input.
func (r *Rectangle) SetSideFromKeyboard(sideName string) float64 {
	fmt.Printf("Enter %s side: ", sideName)
	number, err := keyboard.GetFloat()
	if err != nil {
		return 0
	}
	return number
}

//SetSides sets sizes of rectangle sides.
func (r *Rectangle) SetSides() {
	r.A = r.SetSideFromKeyboard("A")
	r.B = r.SetSideFromKeyboard("B")
	err := r.IsRectangle()
	if err != nil {
		log.Fatal(err)
		return
	}
}

//GetArea returns the area of rectangle.
func (r *Rectangle) GetArea() float64 {
	return r.A * r.B
}

//GetPerimeter returns the perimeter of rectangle.
func (r *Rectangle) GetPerimeter() float64 {
	return (r.A + r.B) * 2
}

func main() {
	var ABC Rectangle
	ABC.SetSides()
	fmt.Println(ABC.GetPerimeter())
	fmt.Println(ABC.GetArea())
}
