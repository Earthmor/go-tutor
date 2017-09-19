package main

import "fmt"

func test(){
	defer fmt.Println("t1")
	fmt.Println("2")
}

func main(){
	test()
	test()
}