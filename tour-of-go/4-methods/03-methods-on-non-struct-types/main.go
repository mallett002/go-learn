package main

import (
    "fmt"
    "math"
)

type MyFloat float64

// note: you can only declare method with receiver whos type is defined in same package
// non-struct type "MyFloat" has Abs function on it (Abs has receiver of type MyFloat)
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }

    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt2)

    fmt.Println(f.Abs())
}
