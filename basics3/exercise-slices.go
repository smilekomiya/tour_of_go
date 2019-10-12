package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slicex := make([]uint8, dx)
	slicey := make([]uint8, dy)
	for i := range slicey {
		slicey[i] = slicex
	}

	return slicey
}

func main() {
	pic.Show(Pic)
}
