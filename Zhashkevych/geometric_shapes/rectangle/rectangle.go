package rectangle

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

//Check user input for negative values of side.
func (r *Rectangle) Check() error {
	if r.A < 0 || r.B < 0 {
		return errors.New("Side size can't be less than zero!")
	}
	return nil
}

//SetSidesFromKeyboard sets sizes of rectangle sides from user input.
func (r *Rectangle) SetSizeFromKeyboard(sideName string) float64 {
	fmt.Printf("Enter %s side: ", sideName)
	number, err := keyboard.GetFloat()
	if err != nil {
		return 0
	}
	return number
}

//SetSize sets sizes of rectangle sides.
func (r *Rectangle) SetSize() {
	r.A = r.SetSizeFromKeyboard("A")
	r.B = r.SetSizeFromKeyboard("B")
	err := r.Check()
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

//Show prints result of calculating area and perimeter of rectangle.
func (r *Rectangle) Show() {
	rArea := c.GetArea()
	rPerimeter := c.GetPerimeter()
	fmt.Printf("Rectangle area: %f cm^2", rArea)
	fmt.Println()
	fmt.Printf("Circumference of the rectangle: %f cm", rPerimeter)
}
