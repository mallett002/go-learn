package main

import (
    "fmt"
    "time"
)

// switch, hits first case and exits
// if hits first case "today + 0", does not go into "today + 1"

func main() {
    fmt.Println("When is Saturday?")
    
    today := time.Now().Weekday()
    tomorrow := today + 1

    fmt.Printf("Today is %v\n", today)
    fmt.Printf("Tomorrow is %v\n", tomorrow)

    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
}
