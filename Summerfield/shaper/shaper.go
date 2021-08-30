package main

import (
	"image"
	"image/color"
)

func main() {
	img := FilledImage(420, 220, image.White)
	fill := color.RGBA{200, 200, 200, 0xFF}
	rectangle := NewRectangle(fill, image.Rect(0, 0, 420, 220), true)
	rectangle.SetFilled(true)
	rectangle.Draw(img)
	SaveImage(img, "rectangle.png")
}
