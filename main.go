package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	v1 := Vertex{1, 2}
	v2 := Vertex{3, 4}
	fmt.Println(v1, v2)
	fmt.Println(v1.X, v2.X)
	fmt.Println(v2.Y, v1.Y)

	// Struct pointer
	v3 := &v1
	v3.X, v3.Y = 10, 20
	fmt.Println(v1)

	// struct literals
	structLiterals()
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)

	p.X = 1000
	fmt.Println(v1, v2, v3, *p)
}
