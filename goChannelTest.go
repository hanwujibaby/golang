package main

import (
	"fmt"
)

func main() {
	//make的第二个参数用于指定channel的缓冲区大小,超过缓冲区的话会报错
	var c = make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
