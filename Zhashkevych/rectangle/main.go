package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Rectangle-struct describes geometric shape rectangle.
type Rectangle struct {
	A int
	B int
}

//isRectangle check user input for negative values of side.
func (r *Rectangle) isRectangle() error {
	if r.A < 0 || r.B < 0 {
		return errors.New("Side size can't be less than zero!")
	}
	return nil
}

//setSidesFromKeyboard sets sizes of rectangle sides from user input.
func (r *Rectangle) setSideFromKeyboard(sideName string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s side: ", sideName)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		return 0
	}
	return int(number)
}

//setSidesFromKeyboard sets sizes of rectangle sides.
func (r *Rectangle) setSides() {
	r.A = r.setSideFromKeyboard("A")
	r.B = r.setSideFromKeyboard("B")
	err := r.isRectangle()
	if err != nil {
		log.Fatal(err)
		return
	}
}

//gerArea returns the area of rectangle.
func (r *Rectangle) getArea() int {
	return r.A * r.B
}

//getPerimeter returns the perimeter of rectangle.
func (r *Rectangle) getPerimeter() int {
	return (r.A + r.B) * 2
}

func main() {
	var ABC Rectangle
	ABC.setSides()
	fmt.Println(ABC.getPerimeter())
	fmt.Println(ABC.getArea())
}
