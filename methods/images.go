package main

import (
	"fmt"
	"image"

	"golang.org/x/tour/pic"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(m.ColorModel())
	pic.ShowImage(m)
}
