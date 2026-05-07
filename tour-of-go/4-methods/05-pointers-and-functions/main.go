package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// normal function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// normal function
// scales the real thing, not just copy (value at memory address)
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}

    Scale(&v, 10)

    fmt.Println(Abs(v))
}
