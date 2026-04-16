package main

import "fmt"


// Defer statment defers execution of something until the surrounding function returns.
// Although, the defered call's args are evaluated immediately.

func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
