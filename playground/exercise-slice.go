package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dx)
	x := 0
	for ; x < dx; x++ {
		result[x] = make([]uint8, dy)
		y := 0
		for ; y < dy; y++ {
			result[x][y] = uint8((x * y) / 2)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
