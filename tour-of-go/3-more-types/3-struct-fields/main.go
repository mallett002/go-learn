package main

import "fmt"

// Struct fields accessible using a dot

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// Reassign X
	v.X = 4

	fmt.Println(v)
}
