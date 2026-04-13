package main

import "fmt"

func main() {
    sum := 1 

    // "for" is go's while loop
    // same as 2-for-loop-omit/main.go but without semi colons
    for sum < 1000 {
        sum += sum
    }

    // can even loop forever:
    // for {} // not sure why you'd want to

    fmt.Println(sum)
}
