package main

import "fmt"

func fibonacci() func() int {
	f, s := 0, 1
	return func() (r int ){
		r = f
		f, s = s, f+s
		return
	}
}

func main() {
	f:=fibonacci()
	for i:=0;i<10;i++{
		fmt.Println(f())
	}
}