package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("Done working")

	done <- true
}

func main() {
	done := make(chan bool, 1) // only have 1 thing in at a time

	go worker(done)

	fmt.Println("Waiting for work to be done...")
	<-done // hangs here until worker is done

	fmt.Println("Process finished")
}
