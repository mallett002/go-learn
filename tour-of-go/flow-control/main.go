package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// basicForLoop()
	// ifStatments()
	// switchStatments()
	deferFunctions()
}

func basicForLoop() {
	sum := 1
	for i := 0; i < 2; i++ {
		sum += i
	}

	// init and post statements are optional:
	for sum < 4 {
		sum += sum
	}

	// for is Go's "while" loop
	for sum < 100 {
		fmt.Printf("\n adding %v to %v", sum, sum)
		sum += sum
	}

	fmt.Println(sum)

	// infinite loop:
	// for {
	// 	fmt.Println("looping...")
	// }
}

func sqrt(x float64) string {
	fmt.Println(x)
	if x < 0 {
		fmt.Printf("\n%v is less than 0\n", x)
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {

	// if can have statement to execute before condition:
	if v := math.Pow(x, n); v < lim {
		return v
	} else { // vars inside if are available inside else blocks:
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// can't use v here though (outside if and elses)
	return lim
}

func ifStatments() {
	// fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func switchStatments() {
	// no break keyword needed

	// Basic switch:
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	// switch without condition (`switch true`)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferFunctions() {
	// defers are pushed onto a stack. LIFO order
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
