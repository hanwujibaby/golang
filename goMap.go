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
	m = make(map[string]MapStruct) //map的创建必须用make,不是new,一个值为nil是不能赋值的
	m["1"] = MapStruct{2, 3}
	fmt.Println(m["1"])
}
