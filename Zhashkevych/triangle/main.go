package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//Triangle-struct describes geometric shape triangle.
type Triangle struct {
	A int
	B int
	C int
}

//isTriangle uses triangle inequality to check user input and determine whether it is a triangle or not.
func (t *Triangle) isTriangle() error {
	if t.A < 0 || t.B < 0 || t.C < 0 {
		return errors.New("Side size can't be less than zero!")
	}
	if t.A > t.B+t.C || t.B > t.C+t.A || t.C > t.A+t.B {
		return errors.New("This is not a triangle!")
	}
	return nil
}

//setSidesFromKeyboard sets sizes of triangle sides from user input.
func (t *Triangle) setSideFromKeyboard(sideName string) int {
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

//setSidesFromKeyboard sets sizes of triangle sides.
func (t *Triangle) setSides() {
	t.A = t.setSideFromKeyboard("A")
	t.B = t.setSideFromKeyboard("B")
	t.C = t.setSideFromKeyboard("C")
	err := t.isTriangle()
	if err != nil {
		log.Fatal(err)
		return
	}
}

//gerArea returns the area of triangle using the Heron's formula.
func (t *Triangle) getArea() float64 {
	sP := float64(t.getPerimeter()) / 2.0
	fmt.Println(sP)
	return math.Sqrt(sP * (sP - float64(t.A)) * (sP - float64(t.B)) * (sP - float64(t.C)))
}

//getPerimeter returns the perimeter of triangle.
func (t *Triangle) getPerimeter() int {
	return t.A + t.B + t.C
}

func main() {
	var ABC Triangle
	ABC.setSides()
	fmt.Println(ABC.getPerimeter())
	fmt.Println(ABC.getArea())
}
