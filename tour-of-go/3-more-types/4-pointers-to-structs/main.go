package main

import "fmt"

// Struct fields can be accessed through a pointer

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	p := &v // pointer to v

	// could do this (but cumbersome to write)
	(*p).X = 1e9

	// go permits us to just do this:
	p.X = 1e9

	fmt.Println(v)
}
