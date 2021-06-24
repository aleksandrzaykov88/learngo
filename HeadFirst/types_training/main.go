package main

import (
	"fmt"
)

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

func (l Liters) toMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func (m Milliliters) toGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (m Milliliters) toLiters() Liters {
	return Liters(m / 1000)
}

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) toMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
}
