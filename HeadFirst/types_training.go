package main

import (
	"fmt"
)

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) toGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) toMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (l Liters) toMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) toLiters() Liters {
	return Liters(m / 1000)
}

func liquidVolumes() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.toGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.toGallons())
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.toLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.toMilliliters())
	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters\n", l, l.toMilliliters())
	ml := Milliliters(500)
	fmt.Printf("%0.1f milliliters is %0.1f liters\n", ml, ml.toLiters())
}
