package main

import (
    "fmt"
    "math"
)

// if can start with short statement to execute before if-check

func pow(num, exp, limit float64) float64 {

    // statement before if:
    if v := math.Pow(num, exp); v < limit {
        // v only in scope here
        return v
    }

    return limit
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
