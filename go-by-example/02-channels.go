package main

import (
	"fmt"
	"time"
)

func main2() {
	msgs := make(chan string)

	go func() {
		time.Sleep(time.Second)

		msgs <- "ping!"
	}()

	msg := <-msgs // waits (blocks) until message is available

	fmt.Println(msg)
}
