package main

import (
    "fmt"
    "math"
)

// if statement in go (no parens); needs "{}"

func sqrt(num float64) string {
    fmt.Println(num)

    if num < 0 {
        return sqrt(-num) + "i"
    }

    return fmt.Sprint(math.Sqrt(num))
}

func main() {
    fmt.Println("answer:", sqrt(2))
    fmt.Println("answer:", sqrt(-4))
}
