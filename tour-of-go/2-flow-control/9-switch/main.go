package main

import (
    "fmt"
    "runtime"
)

// go switch has break built in
// value doesn't need to be a constant
// value doesn't need to be a number

func main() {
    fmt.Print("Go runs on ")

    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("macOS.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Printf("%s.\n", os)
    }
}
