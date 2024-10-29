package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := range ret {
		ret[i] = make([]uint8, dx)
	}
	for i := range ret {
		for j := range ret[i] {
			ret[i][j] = uint8((i+j)/2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
