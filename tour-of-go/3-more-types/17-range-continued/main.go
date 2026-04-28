package main

import "fmt" 


func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // 1, 10, 100, 1000 (binary). Decimal: 1, 2, 4, 8
	}

	for _, val := range pow {
		fmt.Printf("%d\n", val)
	}
}

