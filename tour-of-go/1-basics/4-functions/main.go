package main

import (
	"fmt"
)

// types come after var name
// return type (int)
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(32, 14))
}
