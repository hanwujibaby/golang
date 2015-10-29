// goInterface
package main

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	x, y float64
}

type Rec struct {
	x, y int
}

func (r *Rec) Abs() int {
	return r.x + r.y
}

func (v *Vertex) Abs() float64 {
	return v.x * v.y
}

func main() {
	var a Abser
	x := &Vertex{3.0, 2.0}
	r := &Rec{1, 2}
	v := Vertex{1, 2}
	fmt.Println(v)

	a = x
	//a = &x
	//a = r
	fmt.Println(x.Abs())
	fmt.Println(r.Abs())
	fmt.Println(a)

}
