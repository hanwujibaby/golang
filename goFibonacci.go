// goFibonacci
package main

import (
	"fmt"
)

func fibbonacci() func(x int) int {
	var sum = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	f := fibbonacci()
	for i := 0; i < 6; i++ {
		fmt.Printf("index:%d,value:%d\n", i, f(i))
	}

}
