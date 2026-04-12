package main

import "fmt"

// variables without initial explicit value use their zero values:
// 0 for numeric, false for booleans etc.
// "" for strings

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v, %v, %v, %q", i, f, b, s)
}
