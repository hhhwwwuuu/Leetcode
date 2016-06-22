//map键值与C语言一致
//删除的方式调用delete函数进行删除。

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("the value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
