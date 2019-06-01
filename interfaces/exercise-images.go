package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	w,h := 110,110
	return image.Rect(0, 0, w, h)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x % (y + 1))
	return color.RGBA{v, v, 110, 110}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}