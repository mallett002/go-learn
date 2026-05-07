package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// if receiver not pointer, v is a copy (value, not reference)
// making it a pointer passes by reference and updates the real source
// which is the value at the memory address
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}

    fmt.Printf("before scale: %f\n", v.Abs())

    v.Scale(10)

    fmt.Printf("after scale: %f\n", v.Abs())
}
