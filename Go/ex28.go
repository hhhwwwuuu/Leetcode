//结构体中使用指针

package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

var (
	p = Vertex{1, 2}  //has type vertex
	q = &Vertex{1, 2} //has type *vertex
	r = Vertex{x: 1}  //x:0 and y:0
	s = Vertex{}
)

func main() {
	fmt.Println(p, q, r, s)
}
