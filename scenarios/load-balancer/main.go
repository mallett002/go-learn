package main

import (
	"container/heap"
	"time"
)

// https://go.dev/talks/2012/waza.slide#45

type Request struct {
    fn func() int   // The operation to perform
    c chan int      // The channel to return the result
}

type Worker struct {
    requests    chan Request    // work to do (buffered channel)
    pending     int             // count of pending tasks
    index       int             // index in the heap
}

type Pool []*Worker

type Balancer struct {
    pool Pool
    done chan *Worker
}

func furtherProcess(item int) int {
   return item 
}

// load generator:
// Takes in a write-only work channel for Requests
// Adds work (request) into work channel i.e. generates queue
func requestor(work chan<- Request) {
    c := make(chan int)

    for {
        // kill some time (fake load)
        time.Sleep(time.Second * 2)

        work <- Request{} // send request into work channel

        result := <- c

        furtherProcess(result)
    }
}

// On Worker struct, We have work function that takes in a done channel
func (w *Worker) work(done chan *Worker) {
    // could run this loop body as a goroutine for parallelism
    for {
        req := <- w.requests    // get request from balancer (queue)
        req.c <- req.fn()       // do work and send result to result channel
        done <- w               // we've finished this request
    }
}

// Balancer function
func (b *Balancer) balance(work chan Request) {
    for {
        select {
        case req := <-work: // received a request...
            b.dispatch(req) // ...so, send it to a worker
        case w := <-b.done: // a worker has finished...
            b.completed(w) // ...so, update its info
        }
    }
}

// make Pool an implemenetation of Heap interface (by providing Less method)
// balance by making the Pool a heap tracked by load
func (p Pool) Less(i, j int) bool {
    return p[i].pending < p[j].pending
}

func (p Pool) Push(w *Worker) {
    // return p[i].pending < p[j].pending
    *p = append(*p, w.(*Worker))
}

// Send request to worker
func (b *Balancer) dispatch(req Request) {
    // Grab the least loaded worker
    w := heap.Pop(&b.pool).(*Worker)

    // send it the task
    w.requests <- req

    // increment it's work queue
    w.pending++

    // Put it into its place on the heap
    heap.Push(&b.pool, w)
}

// Completed
func (b *Balancer) completed(w *Worker) {
    // one fewer in the queue
    w.pending--

    // Remove it from the heap
    heap.Remove(&b.pool, w.index)

    // Put it into its place on the heap
    heap.Push(&b.pool, w)
}

