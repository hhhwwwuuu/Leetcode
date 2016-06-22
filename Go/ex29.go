package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

func main() {
	v := new(Vertex)
	fmt.Println(*v)
	v.x, v.y = 11, 9
	fmt.Println(*v)
}
