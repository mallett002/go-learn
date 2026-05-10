package main

import (
	"fmt"
)

type Animal interface {
    Speak() string
}

type Bear struct {
    Sound string
}

// Bear implements Animal here.
// We don't need to explicitly declare that it does so.
func (b Bear) Speak() string {
    return b.Sound
}

func main() {
    var bear Animal = Bear{"roar"}    

    fmt.Println(bear.Speak())
}

