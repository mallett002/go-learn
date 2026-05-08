package main

import (
	"fmt"
	"math"
)

// interface: set of method signatures
// A value of type "interface" can hold any value that implements those methods


// Any type with an `Abs() float64` method satisfies the Abser interface implicity
type Abser interface {
    Abs() float64
}

type MyFloat float64

type Vertex struct {
    X, Y float64
}


// Type MyFloat implemenets Abser interface
// Abs does one thing for a MyFloat type
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }

    return float64(f)
}

// Type MyFloat implemenets Abser interface
// Abs does something else for a *Vertex type
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    // Create the interface container
    var abser Abser

    myFloat := MyFloat(-math.Sqrt2)
    vertex := Vertex{3, 4}

    // a MyFloat implements Abser
    abser = myFloat
    fmt.Println("abser:", abser)
    fmt.Println("abs():", abser.Abs())

    fmt.Println()

    // a Vertex implements Abser
    abser = &vertex
    fmt.Println("abser:", abser)

    // does not compile:
    // abser = vertex // (mandatory that it's a *Vertex type)
    
    fmt.Println("abs():", abser.Abs())
}
