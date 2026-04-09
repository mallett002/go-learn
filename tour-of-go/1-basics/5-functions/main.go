package main

import (
	"fmt"
)

// x & y have same type, can omit all but last one's type:
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(32, 13))
}
