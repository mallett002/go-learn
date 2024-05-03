package main

import (
	"fmt"
	"time"
)

func main() {
	// basics()
	// withGoRoutines()
	// withLoops()
	withBufferedChannel()
}

func basics() {
	var c = make(chan int) // unbuffered channel (holds only 1 val)
	c <- 1                 // blocks here until something reads from it (can't make it to next line)
	var i = <-c
	fmt.Println(i)
}

func withGoRoutines() {
	var c = make(chan int)

	go process(c)

	fmt.Println(<-c)
}

func withLoops() {
	var c = make(chan int)

	// 1 call process
	go processWithLoop(c)

	// 2 loop over the channel (waits for things to start coming into channel)
	for i := range c {
		// pull values from channel
		fmt.Println(i)
	}
}

func withBufferedChannel() {
	var c = make(chan int, 5) // buffered channel

	go processWithLoop(c) // calls here, doesn't hang around (exits before pulling stuff from channel below) Due to it being buffered channel (fits more than 1 val)

	for i := range c {
		time.Sleep(time.Second * 1) // sleep for 1 sec (do some work that takes about a second)
		fmt.Println(i)
	}
}

func process(c chan int) {
	c <- 123 // put 123 into channel
}

func processWithLoop(c chan int) {
	defer close(c) // do this right before function exec runs (notify other processes using the channel that it's done)

	// 3 starts putting things in channel
	for i := 0; i < 5; i++ {
		c <- i
	}

	fmt.Println("Exiting process")
}
