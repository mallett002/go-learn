package main

import "fmt"


// Empty interface is one that specifies zero methods
// Every type has at least 0 methods on it, so everything implements the empty interface

func main() {
    var i interface{}
    describe(i)

    i = 42
    describe(i)

    i = "hello"
    describe(i)
}

func describe(i interface{}) {
    fmt.Printf("( %v, %T )\n", i, i)
}

