package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.x = 1e9
	fmt.Println(p)
}
