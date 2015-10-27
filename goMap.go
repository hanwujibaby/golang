// goMap
package main

import (
	"fmt"
)

//go Map学习

type MapStruct struct {
	x int
	y int
}

var m map[string]MapStruct

func main() {
	m = make(map[string]MapStruct)
	m["1"] = MapStruct{2, 3}
	fmt.Println(m["1"])
}
