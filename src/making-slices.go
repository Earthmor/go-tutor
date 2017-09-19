package main

import "./utils"

func main(){
	a := make([]int, 5)
	utils.PrintStrSlice("a", a)

	b := make([]int, 0, 5)
	utils.PrintStrSlice("b", b)

	c := b[:2]
	utils.PrintStrSlice("c", c)

	d := c[2:5]
	utils.PrintStrSlice("d", d)
}
