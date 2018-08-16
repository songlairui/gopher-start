package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (o Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, o.w, o.h)
}

func (o Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (o Image) At(x int, y int) color.Color {
	return color.RGBA{11, 222, 255, 255}
}

func main() {
	m := Image{10, 10}
	//fmt.Println(m)
	pic.ShowImage(m)
}
