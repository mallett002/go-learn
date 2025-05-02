package main

import (
	"fmt"
	"time"
)

func main3() {
	msgs := make(chan string, 2)

	fmt.Println("Sending 1")
	msgs <- "1"
	fmt.Println("Sent 1")

	fmt.Println("Sending 2")
	msgs <- "2"
	fmt.Println("Sent 2")

	// Sending 3 (which will block due to not enough room in channel)
	fmt.Println("Sending 3")
	go func() {
		msgs <- "3"
		fmt.Println("Sent 3")
	}()

	fmt.Println("Waiting...")
	time.Sleep(2 * time.Second)

	// pull from channel (allow "3" to send by adding room to channel)
	val := <-msgs
	fmt.Println("received: ", val)

	// give "3" a chance to send
	time.Sleep(time.Second)
}
