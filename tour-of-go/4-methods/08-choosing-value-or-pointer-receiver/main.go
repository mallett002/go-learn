package main

import (
	"fmt"
	"math"
)

// 2 main reasons to use pointer receivers:
//   - modify underlying values
//   - avoid copying data for memory efficiency
// all methods should be either pointer or non-pointers. don't mix them on a given type

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// vertex as a non-pointer value
	v := Vertex{3, 4}

	fmt.Printf("Before scaling: %+v; Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v; Abs: %v\n", v, v.Abs())
}
