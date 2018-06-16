package main

import "golang.org/x/tour/pic"
func Pic(dx, dy int) [][]uint8 {
	var images = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		images[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			images[i][j] = uint8( i * j)
		}
	}
	return images
}

func main() {
	pic.Show(Pic)
}