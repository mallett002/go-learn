package main

import "fmt"


// constants declared with const keyword
// cannot use ":=" with constants

// Pascal case by convention (if you want them exported)
// Only exported in global scope
const Pi = 3.14

func main() {
	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("", Truth)
}
