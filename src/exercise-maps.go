package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int{
	r := make(map[string]int)
	f := strings.Fields(s)
	for _, v := range f {
		if r[v] == 0 {
			r[v] = 1
		} else {
			r[v] = r[v] + 1
		}
	}
	return r
}

func main(){
	//wc.Test(WordCount) just we don't have wc package
	fmt.Println("result is : ", WordCount("I I I I am learning Go!"))
}
