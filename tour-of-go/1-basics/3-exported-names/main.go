package main

import (
	"fmt"
	"math"
)

// exported items from a package start with a capital letter (such as math.Pi)
// lowercase items aren't accessible outside the package
func main() {
	fmt.Println("Pi is", math.Pi)
}
