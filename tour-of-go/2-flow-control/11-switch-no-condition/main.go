package main

import (
    "fmt"
    "time"
)

// switch, doesn't need condition
// clean way to write long if/else statements

func main() {
    t := time.Now()

    switch {
    case t.Hour() < 12:
        fmt.Println("Good Morning!")
    case t.Hour() < 17:
        fmt.Println("Good Afternoon.")
    default:
        fmt.Println("Good Evening.")
    }
}
