package main

import (
	"fmt"
	"math/cmplx"
)

// variables factored into block
var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1 // shifts 1 left 64 times in binary: 10000000000.. 64 0s (then subtracts 1)

	// i is imaginary (allows you to do square roots of negative numbers:
	//  √(-4) = 2i because (2i)² = 4i² = 4(-1) = -4
	z complex128 = cmplx.Sqrt(-5 + 12i) // (2+3i)
)

/*
types:
	- int int8 int16 int32 int64 (use int unless good reason not to)
	- uint uint8 uint16 uint32 uint64 uintptr (32 size on 32-bit systesm and 64 on 64-bit)
	- byte alias for uint8
	- rune alias for int32 (unicode code point)
	- float32 float64
	- complex64 complex128

	- just use int, uint and float64 unless good reason not to
		(e.g., memory constraints in large arrays, matching specific API)
*/

func main() {
	fmt.Printf("Type: %T; Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T; Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T; Value: %v\n", z, z)
}


