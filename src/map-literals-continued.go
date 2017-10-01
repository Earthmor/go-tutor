package main

import (
	"fmt"
	"./utils"
)

var m = map[string]utils.Vertex{
	"Bell Labs": {40.68433,-74.39967},
	"Google": {37.42202, -122.08408},
}

func main(){
	fmt.Println(m)
}