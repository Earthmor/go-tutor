package main

import (
	"./utils"
)

func main(){
	var s []int
	utils.PrintSlice(s)

	s = append(s, 0)
	utils.PrintSlice(s)

	s = append(s, 1)
	utils.PrintSlice(s)

	s = append(s, 2, 3, 4)
	utils.PrintSlice(s)
}