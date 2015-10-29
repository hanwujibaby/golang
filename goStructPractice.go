// goStructPractice
package main

import (
	"fmt"
)

type Rec struct {
	width  int
	height int
}

func (r *Rec) area() int {
	r.width = r.width * 2
	r.height = r.height * 2
	return r.width * r.height
}

func main() {
	x := new(Rec)
	x.height = 2
	x.width = 3
	fmt.Println(x.area())
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println("Hello World!")
}
