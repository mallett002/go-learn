package main

import "fmt"

func main() {
	m := make(map[string]int)

	// add/update value at key:
	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	m["answer"] = 48
	fmt.Println("The value:", m["answer"])

	// delete a key-value pair
	delete(m, "answer")
	fmt.Println("The value:", m["answer"])

	// check if present
	v, ok := m["answer"]
	fmt.Println("The value:", v, "Is present?", ok)
}
