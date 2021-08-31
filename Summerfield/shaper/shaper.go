package main

import (
	"image"
	"image/color"
)

func main() {
	img := FilledImage(420, 220, image.White)
	fill := color.RGBA{200, 200, 200, 0xFF}
	for i := 0; i < 10; i++ {
		width, height := 40+(20*i), 20+(10*i)
		rectangle := NewRectangle(fill, image.Rect(0, 0, width, height), true)
		rectangle.SetFilled(true)
		rectangle.Draw(img, i*10, i*10)
		fill.R -= uint8(i * 5)
		fill.G = fill.R
		fill.B = fill.R
	}
	SaveImage(img, "rectangle.png")
}
