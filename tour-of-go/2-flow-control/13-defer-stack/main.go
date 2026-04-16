package main

import "fmt"


// Defer statements are pushed onto a stack
// They are executed in LIFO order (stack)

func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Printf("numero %v\n", i)
    }

    fmt.Println("done")
}

// ouptut:
// counting
// done
// numero 9
// numero 8
// numero 7
// numero 6
// numero 5
// numero 4
// numero 3
// numero 2
// numero 1
// numero 0
