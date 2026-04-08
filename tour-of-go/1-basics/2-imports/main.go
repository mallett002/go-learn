package main

// import with factored import statement is best practice:
import (
	"fmt"
	"math"
)

// Could also do this
// import "fmt"
// import "math/rand"

func main () {
	fmt.Printf("Now you have %g problems \n", math.Sqrt(7))
}
