package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

//Bounds is an implementation of basic image-interface method.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 150, 150)
}

//ColorModel is an implementation of basic image-interface method.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

//At is an implementation of basic image-interface method.
func (i Image) At(x, y int) color.Color {
	return color.RGBA{100, 200, 200, 0xff}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
