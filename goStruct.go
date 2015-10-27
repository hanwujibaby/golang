// goStruct
package main

import (
	"fmt"
)

type TestStruct struct {
	x int
	y int
}

func main() {
	p := new(TestStruct)
	p.x, p.y = 20, 10
	fmt.Println(p)

}
