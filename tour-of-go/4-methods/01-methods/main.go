package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// define methods on types
// Abs lives on Vertex
// Abs has receiver of type Vertex
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
