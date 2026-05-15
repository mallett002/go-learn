package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

// implements error interface
func (e ErrNegativeSqrt) Error() string {
    // Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
    // You can avoid this by converting e first: fmt.Sprint(float64(e)).
    // It causes the infinite loop because fmt package checks if "e" implements error
    // If it implmentes error, it calls the Error() method.
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return x, ErrNegativeSqrt(x)
    }

    const OneToPowerNegTen = 1e-10

    z := float64(1)
    lastVal := float64(10)
    diff := math.Inf(1) // positive infinity (sign is 1 and not -1)

    for diff > OneToPowerNegTen {
        z -= (z * z - x) / (2 * z)
        diff = math.Abs(z - lastVal)
        lastVal = z
    }

    return 0, nil
}

func main() {
    // Println looks for things that implement error (calls Error() method if it's an error)
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
