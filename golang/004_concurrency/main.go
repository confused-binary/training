package main

import (
	"fmt"
	"time"
)

func say(s string) {
	// Because we have no control over order of execution,
	// without syncronization they can execute at any time.
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	// Received data via 'c' channel
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Send results 'c' channel
	// No need to specify return details in function id
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// practice to get a better grasp on golang concurrency
// .... at least that's the hope
func main() {
	// Concurrency is accomplished with go routines.
	// Each go routine operates in the same address space
	// https://go.dev/tour/concurrency/1
	go say("World")
	say("Hello")

	// Channels are typed conduit that data can be sent and received from
	// <- operator is reserved to accomplish this assignment.
	// Data flows in the direction of the arrows
	// https://go.dev/tour/concurrency/2
	s := []int{7, 2, 8, -9, 4, 0}
	// By default, sends and receives data until the other side is ready.
	// Allowing go routines to synchronize without explicit locks
	// First ceate a channel to transfer data with
	c := make(chan int)
	// Then execute functions as go routines
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// Finally receive data from the channels
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	// Channels can be buffered to provide time for downstream
	// tasks to complete actions without missing data
	// If you try to add data to a channel with no space it will deadlock
	// https://go.dev/tour/concurrency/3
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// If needed, channels can be closed, which can be useful
	// if the channel is going to be iterated over as a range
	// Otherwise the channel will deadlock when it can't add more
	// data values than the configured buffer
	// https://go.dev/tour/concurrency/4
	cha := make(chan int, 10)
	go fibonacci(cap(cha), cha)
	for i := range cha {
		fmt.Println(i)
	}

	// Select statement lets a goroutine wait on multiple communication operations.
	// This is useful for 'case' evaluations
	// https://go.dev/tour/concurrency/5
	// The 'default' case can be used so that the go routine doesn't stall
	// https://go.dev/tour/concurrency/6

	fmt.Println("That's it!")
}
