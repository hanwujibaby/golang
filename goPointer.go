// goPointer
package main

import (
	"fmt"
)

func main() {
	//整形变量
	var a int
	//指针变量；存放的是指向int类型的变量内存地址
	var ptr *int

	a = 4
	ptr = &a //&为获取地址的标记
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) //*ptr 表示为获取指针变量的值

}
