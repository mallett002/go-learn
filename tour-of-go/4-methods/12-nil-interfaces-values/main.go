package main

import "fmt"


// A nil interface holds neither value nor concrete type
// There's no type inside the interface typle (value, type), so it doesn't know which method to call

type I interface {
    M()
}

func main() {
    var i I
    
    describe(i)

    // Interface wasn't implemented
    // doesn't know which method to call
    i.M()
}

func describe(i I) {
    fmt.Printf("( %v, %T )\n", i, i)
}

