package main

import (
	"image"
	"image/color"
)

func main() {

}

type Image struct {
	w, h int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 255, 255}
}
