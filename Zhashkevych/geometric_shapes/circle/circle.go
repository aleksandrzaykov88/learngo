package circle

import (
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/aleksandrzaykov88/learngo/HeadFirst/keyboard"
)

const pi = math.Pi

//Circle-struct describes geometric shape circle.
type Circle struct {
	R float64
}

//SetSize sets circle radius.
func (c *Circle) SetSize() {
	c.R = c.SetSizeFromKeyboard("")
}

//SetSizeFromKeyboard sets circle radius from user input.
func (c *Circle) SetSizeFromKeyboard(str string) float64 {
	fmt.Println("Enter a radius of the circle: ", str)
	number, err := keyboard.GetFloat()
	if err != nil {
		return 0
	}
	return number
}

//Show prints result of calculating area of circle.
func (c *Circle) Show() {
	circleArea := c.GetArea()
	circumference := c.GetPerimeter()
	fmt.Printf("Circle radius: %f cm", c.R)
	fmt.Println()
	fmt.Printf("Circle area: %f cm^2", circleArea)
	fmt.Println()
	fmt.Printf("Circumference of the circle: %f cm", circumference)
}

//Check user input for negative values of side.
func (c *Circle) Check() error {
	if c.R < 0 {
		return errors.New("Radius can't be less than zero!")
	}
	return nil
}

//GetArea calculates the area of a circle with a given radius.
func (c *Circle) GetArea() float64 {
	err := c.Check()
	if err != nil {
		log.Fatal(err)
	}
	return c.R * c.R * pi
}

//GetPerimeter calculates the circumference
func (c *Circle) GetPerimeter() float64 {
	err := c.Check()
	if err != nil {
		log.Fatal(err)
	}
	return 2 * c.R * pi
}
