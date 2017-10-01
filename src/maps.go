package main

import (
	"fmt"
	"./utils"
)
var m map[string]utils.Vertex

func main(){
	m = make(map[string]utils.Vertex)
	m["Bell Labs"] = utils.Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
