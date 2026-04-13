package main

import "fmt"


// Numeric constants are high precision values
// Untyped constant takes precision needed
const (
	// Create a huge number
	Big = 1 << 100 // shift 1 left 100 times in binary (1 then 100 zeros after it)
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1	
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
