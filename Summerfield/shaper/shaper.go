package main

import (
	"image"
	"image/color"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	red := color.RGBA{0xFF, 0, 0, 0xFF}
	circ := NewCircle(red, 25)
	circ.Fill()
	circ.Draw(img, 50, 50)
	SaveImage(img, "test.png")
}
