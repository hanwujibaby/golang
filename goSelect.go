package main

import (
	"fmt"
)

func test(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Printf("quit")
			return
		}
	}

}

func main() {
	var c = make(chan int)
	var quit = make(chan int)

	//用go新起一个协程去运行；会等待channel c的输入
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	test(c, quit)

}
