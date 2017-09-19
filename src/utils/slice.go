package utils

import "fmt"

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func PrintStrSlice(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
