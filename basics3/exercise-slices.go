package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slicex := make([]uint8, dx)
	for i := range slicex {
		slicex[i] = uint8(i)
	}

	slicey := make([][]uint8, dy)
	for i := range slicey {
		slicey[i] = slicex
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			slicey[y][x] = uint8((x + y) / 2)
		}
	}

	return slicey
}

func main() {
	pic.Show(Pic)
}
