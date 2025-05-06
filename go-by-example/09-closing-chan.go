package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// pull from jobs and process them
	// notify on done chan when we've received all jobs:
	go func() {
		for {
			j, more := <-jobs // more will be false if jobs has been closed

			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("all jobs completed")
				done <- true
				return
			}
		}
	}()

	// get jobs to work and send them into jobs chan
	// send 3 workers over the jobs chan then close it
	for j := 1; j < 4; j++ {
		jobs <- j
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// await the worker with channel sync
	<-done

	// ensure no more jobs
	_, more := <-jobs
	fmt.Println("received more jobs:", more)
}
