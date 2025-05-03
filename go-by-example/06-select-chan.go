package main

import (
	"fmt"
	"time"
)

func main6() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	// simulate RPC procedures:
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "three"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1:", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2:", msg2)
		case msg3 := <-c3:
			fmt.Println("msg3:", msg3)
		}
	}

	fmt.Println("Done")
}
