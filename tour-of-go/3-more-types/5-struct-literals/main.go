package main

import "fmt"

// A struct literal denotes a newly allocated struct value by listing the values of its fields

type Vertex struct {
	X, Y int
}

// struct literals: combines declaration + initialization with values in 1 expression
var (
	v1 = Vertex{1, 2} // has type vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{} // X:0 and Y:0
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
