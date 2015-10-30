package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Printf("string %s,index %d\n", s, i)
	}
}

func main() {
	//go 关键字，相当于启动java一个线程运行say方法
	go say("world")
	say("hello")
}
