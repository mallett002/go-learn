package main

import (
    "fmt" 
    "math"
)


type Vertex struct {
	X, Y float64
}

// Non-Pointer receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer arg
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    // vertex as a non-pointer value
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
    fmt.Println(AbsFunc(v))

    // vertex as a pointer value
    p := &Vertex{4, 3}
    // Even though receiver requires non-pointer, works also with pointer:
    fmt.Println(p.Abs())
    // function does not work same way, have to pass the actual type:
    fmt.Println(AbsFunc(*p))
}
