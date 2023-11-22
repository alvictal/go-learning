package main

/* Work with thread on Go it is a bit interesting and simples
follow the example below to understand how threads work in go */

import (
	"fmt"
	"time"
)

func main() {
	// Create a piece of memmory that could shared for other threads
	channel := make(chan int)

	// create the thread with the go statment and a anonnymus functions
	go func() {
		for i := 0; i <= 1000; i++ {
			//Put on the channel - that is the shared of piece shared to other threads
			channel <- i
			fmt.Printf("Put on channel value %v Thread 1 executing\n", i)
		}
	}()

	go func() {
		for i := 1001; i <= 2000; i++ {
			channel <- i
			fmt.Printf("Put on channel value %v Thread 2 executing\n", i)
		}
	}()

	go worker(channel, 1)
	go worker(channel, 2)
	go worker(channel, 3)
	worker(channel, 4)

}

func worker(channel chan int, workerID int) {
	for {
		// Remove from the channel the value that was storage there
		fmt.Println("Pick up from chanel", <-channel, "workerId", workerID)
		time.Sleep((time.Second))
	}
}
