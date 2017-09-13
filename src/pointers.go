package main

import "fmt"

func main() {
	i, j := 40, 1234
	p := &i
	fmt.Println(*p)
	*p = 20
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}