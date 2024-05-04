package main

import (
	"fmt"
)

func main() {
	// 3 slices, for int float32 and float64
	var intSlice = []int{1, 2, 3, 4}
	// fmt.Println(sumIntSlice(intSlice))
	// instead of ^^ use the generic function:
	fmt.Println(sumSlice[int](intSlice))
	// can remove "[T]", it is inferred:
	sumSlice(intSlice)
	
	var float32Slice = []float32{1, 2, 3, 4}
	// fmt.Println(sumFloat32Slice(float32Slice))
	// instead of ^^ use the generic function:
	fmt.Println(sumSlice[float32](float32Slice))
	
	var float64Slice = []float64{1, 2, 3, 4}
	// fmt.Println(sumFloat64Slice(float64Slice))
	// instead of ^^ use the generic function:
	fmt.Println(sumSlice[float64](float64Slice))

	fmt.Printf("\nis empty: %v", isEmpty[int](intSlice))
	fmt.Printf("\nis empty: %v", isEmpty[int]([]int{}))
}

// instead of the 3 functions below, use a generic function
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T

	for _, v := range slice {
		sum += v
	}

	return sum
}

func sumIntSlice(slice []int) int {
	var sum int

	for _, v := range slice {
		sum += v
	}

	return sum
}

func sumFloat32Slice(slice []float32) float32 {
	var sum float32

	for _, v := range slice {
		sum += v
	}

	return sum
}

func sumFloat64Slice(slice []float64) float64 {
	var sum float64

	for _, v := range slice {
		sum += v
	}

	return sum
}

// useing "any" type
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
