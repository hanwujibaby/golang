// goMap1
package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]int)
	m["1"] = 1
	//修改map值
	m["1"] = 2

	delete(m, "2") //删除map中的元素
	v, ok := m["1"]
	fmt.Printf("s")
	fmt.Println(v, ok)
}
