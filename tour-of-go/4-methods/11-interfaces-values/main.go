package main

import (
	"fmt"
	"math"
)

// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Each type can have its own implmentation of how the method should work

type I interface {
    M()
}

type T struct {
    S string
}

type F float64

// T implements I
func (t *T) M() {
    fmt.Println(t.S)
}

// T implements I
func (f F) M() {
    fmt.Println(f)
}

func describe(i I) {
    fmt.Printf("( %+v, %T )\n", i, i)
}

func main() {
    var i I

    i = &T{"Hello"}
    describe(i)
    // call &T's implementation of M
    i.M()


    i = F(math.Pi)
    describe(i)
    // call F's implementation of M
    i.M()
}

