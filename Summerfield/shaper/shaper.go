package main

import (
	"fmt"
	"image/color"
)

func main() {
	red := color.RGBA{0xFF, 0, 0, 0xFF}
	o := Option{red, 3}
	s, _ := New("hexagon", o)
	fmt.Println(s)
}

func sanityCheck(name string, shape Shaper) {
	fmt.Print("name=", name, " ")
	fmt.Print("fill=", shape.Fill(), " ")
	if shape, ok := shape.(Radiuser); ok {
		fmt.Print("radius=", shape.Radius(), " ")
	}
	if shape, ok := shape.(Sideser); ok {
		fmt.Print("sides=", shape.Sides(), " ")
	}
	fmt.Println()
}

func showShapeDetails(shape Shaper) {
	fmt.Print("fill=", shape.Fill(), " ")
	if shape, ok := shape.(Radiuser); ok {
		fmt.Print("radius=", shape.Radius(), " ")
	}
	if shape, ok := shape.(Sideser); ok {
		fmt.Print("sides=", shape.Sides(), " ")
	}
	fmt.Println()
}
