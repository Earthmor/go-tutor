package main

import "fmt"

func pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8(i*j)
		}
	}
	return pic
}

func picR(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8(y*x)
		}
	}
	return pic
}

func main(){
	//it is fake, because there isn't package for making picture
	result := pic(12,23)
	fmt.Println(result)

	resultR := picR(1,10)
	fmt.Println(resultR)
}