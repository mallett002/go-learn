package main

import "fmt"


// If concrete value of interface is nil, receiver method will be called with nil
// Note: an interface value that holds the concrete value is itself non-nil

type I interface {
    M()
}

type T struct {
    S string
}

// receiver type can be nil here
func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }

    fmt.Println(t.S)
}

func describe(i I) {
    fmt.Printf("( %v, %T )\n", i, i)
}

func main() {
    var i I
    var t *T
    
    i = t

    describe(i)
    i.M()
}

