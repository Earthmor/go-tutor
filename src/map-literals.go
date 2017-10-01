package main

import (
	"fmt"
	"./utils"
)

var m = map[string]utils.Vertex{
	"Bell Labs": utils.Vertex{
		40.68433, -74.39967,
	},
	"Google": utils.Vertex{
		37.42202, -122.08408,
	},
}

func main(){
	fmt.Println(m)
}