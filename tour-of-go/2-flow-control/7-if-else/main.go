package main

import (
    "fmt"
    "math"
)

func pow(num, exp, limit float64) float64 {
    if v := math.Pow(num, exp); v < limit {
        // v in scope here
        return v
    } else { // v also in scope here
        fmt.Printf("%g >= %g\n", v, limit)
    }

    // v not in scope here
    return limit
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
