package main

import (
	"fmt"
	"math"
)

// T(v) converts v to type T

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, f, z)
	fmt.Printf("Type: %T; Value: %f\n", f, f)
}
