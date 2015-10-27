// goSlice
package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //[]T 表示为数值类型为T的数组
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	//重新切片
	fmt.Println("array[1:4]===>", array[0:4])

	//array[0:0]是没有数据的
	fmt.Println("array[3:3]===>", array[0:0])

	fmt.Println("Hello World!")

	var a = make([]int, 5)
	printSlice("a", a)

	var b = make([]int, 0, 6)

	printSlice("b", b)

	b = b[:cap(b)] //强制通过slice进行扩容
	printSlice("b", b)

	var z []int
	//slice的空值是nil
	if z == nil {
		fmt.Println("nil")
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d,cap=%d,%v\n", s, len(x), cap(x), x)
}
