package main

import (
	"fmt"
)

//c 为一个channel变量；用于与goruntime的通讯,
func sum(a []int, c chan int) {
	var sum = 0
	for _, v := range a {
		sum += v

	}
	c <- sum
}

func main() {
	a := []int{7, 2, 3, 1, 3, 4, -2, 0, -9}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Printf("x:%d,y:%d", x, y)

}
