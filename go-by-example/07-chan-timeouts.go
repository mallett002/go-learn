package main

import (
	"fmt"
	"time"
)

func main7() {
	// Mimic api call. Use non-buffered so non-blocking on the send
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Select with timeout
	// If c2 takes longer than a second, wait no longer
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // need 1 * here?
		fmt.Println("timeout 1")
	}

	// Same demonstration, but with longer timeout that succeeds
	// Mimic api call. Use non-buffered so non-blocking on the send
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// Select with timeout
	// If c2 takes longer than a second, wait no longer
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): // need 1 * here?
		fmt.Println("timeout 2")
	}
}
