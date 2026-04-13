package main

import "fmt"

// init: executed before first iteration
// condition: executed before every iteration
// post: executed end of every iteration

func main() {
    sum := 0 

    for i := 0; i < 10; i++ {
        sum += i
    }

    fmt.Println(sum)
}
