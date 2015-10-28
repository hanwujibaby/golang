// goFunc
package main

import (
	"fmt"
	//"math"
)

//闭包＝函数+引用环境；闭包中的变量相当于引用传递而非值传递
func add() func(int) int {
	var sum = 0
	fmt.Printf("sum:%d;addr of sum:%p\n", sum, &sum) //闭包使用的时候；sum的地址一直没有变化
	return func(x int) int {
		sum += x
		fmt.Printf("x:%d;addr of x:%p\n", x, &x)
		fmt.Printf("sum:%d;addr of sum:%p\n", sum, &sum)
		return sum
	}
}

//函数也是值
func main() {
	//var f = func(x, y float64) float64 {
	//		return math.Sqrt(x*x + y*y)
	//	}

	pos, neg := add(), add()

	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
