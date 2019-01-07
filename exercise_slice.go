package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
	}

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			image[i][j] = uint8((i + j) / 2)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
