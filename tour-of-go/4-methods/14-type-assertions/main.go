package main

import "fmt"

// t := i.(T)
// asserts that i holds type T and assings underlying value of T to t

func main() {
    var i interface{} = "hello"

    // ensure i is a string and set value of s to i's value
    s := i.(string)
    fmt.Println(s) // hello

    s, ok := i.(string)
    fmt.Println(s, ok) // hello, true

    f, ok := i.(float64)
    fmt.Println(f, ok) // 0, false
    // ^^ f becomes 0 value of float64

    // panic:
    // f = i.(float64)
    // fmt.Println(f)
}

