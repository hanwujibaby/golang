// goTour
package main

import "fmt"

//import "math"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

//结构体；类似类？
type Vertex struct {
	X int
	Y int
}

func main() {
	//fmt.Println("my favotite number is:", rand.Intn(10))
	//a, b := swap("hello", "world")
	//fmt.Println(a, b)
	//var i int
	//i := 4
	//fmt.Println(i)
	//常量
	//const isTrue = true
	//fmt.Println(Big * 0.1)

	/*var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	*/

	//s := Vertex{2, 3}
	//fmt.Println(s.X)

	var a int    //整形数据
	var ptr *int //int类型的指针

	a = 3
	ptr = &a
	fmt.Println(*ptr)
}

//好奇怪的语法
func add(x int, y int) int {
	return x + y
}

//函数可以有多个返回值
func swap(a, b string) (string, string) {
	return b, a
}

var c, python, java bool
