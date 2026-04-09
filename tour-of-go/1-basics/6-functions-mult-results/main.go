package main

import (
	"fmt"
)

// func can return multiple results - swap returns 2 strings:
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")

	fmt.Println(a, b)
}
