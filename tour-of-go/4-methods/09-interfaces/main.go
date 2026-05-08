package main

import (
	"fmt"
	"math"
)

// Interface: set of method signatures
// A value of type "interface" can hold any value that implements those methods


// Any type with an `Abs() float64` method satisfies the Abser interface implicity
type Abser interface {
    Abs() float64
}

// MyFloat can implement Abser if has a Abs() method
type MyFloat float64

// Vertex can implement Abser if has a Abs() method
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

// function that accepts an Abser interface
func doubleAbs(a Abser) float64 {
    return a.Abs() * 2
}

func main() {
    // Create the interface container
    var abser Abser // probably don't do this much in practice, just to show that Abser can be any type (MyFloat, *Vertex) as long as they have an Abs method 

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
    fmt.Println("abs():", abser.Abs())

    // does not compile:
    // abser = vertex // (mandatory that it's a *Vertex type)


    // Don't "NEED" the abser wrapper:
    vertex2 := Vertex{5, 6}
    fmt.Println("abs() on vertex2:",  vertex2.Abs())

    fmt.Println("MyFloat doubleAbs: ", doubleAbs(myFloat))
    fmt.Println("*Vertex doubleAbs: ", doubleAbs(&vertex))
}
