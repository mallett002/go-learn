package main

import (
    "fmt"
    "math"
)

// Given x, find number, z, where z^2 is most nearly x
// ex: x: 25
//  z^2 ~ 25
//  5^2 = 25 
//  answer: 5

func Sqrt(x float64) float64 {
    const OneToPowerNegTen = 1e-10
    z := float64(1)
    lastVal := float64(10)
    diff := math.Inf(1) // positive infinity (sign is 1 and not -1)

    for diff > OneToPowerNegTen {
        z -= (z * z - x) / (2 * z)
        diff = math.Abs(z - lastVal)
        lastVal = z
    }

    return z
}


func main() {
    fmt.Println(Sqrt(9))
}
