package main

import (
	"fmt"
	"time"
)

func worker1(ch chan string) {
	time.Sleep(3 * time.Second) // Simulating delay
	for msg := range ch {       // Receive messages until the channel is closed
		fmt.Println("Worker processed:", msg)
	}
}

func main() {
	ch := make(chan string, 3) // Buffered channel with capacity 3

	go worker1(ch) // Start worker Goroutine

	ch <- "pizza is ready"
	ch <- "ice cream is ready"
	ch <- "burger is ready"

	fmt.Println("All orders sent to the buffer")

	time.Sleep(5 * time.Second)
	close(ch) // Close the channel when done
}
