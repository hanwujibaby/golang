// foRange
package main

import (
	"fmt"
	"time"
)

func main() {
	var sum = 0
	var array = []int{1, 2, 4, 5, 6, 8}
	for i, value := range array {
		fmt.Printf("index:%d,value:%d\n", i, value)
	}

	for _, e := range array {
		fmt.Printf("value:%d\n", e)

	}
	time.Sleep(2)
	fmt.Println(sum)
}
