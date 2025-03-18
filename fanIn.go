// Fan-In means multiple goroutines send results to a single channel, allowing one goroutine to aggregate the output.
package main

import (
	"fmt"
	"time"
)

func worker(id int, results chan<- string) {
	for i := 1; i <= 3; i++ {
		results <- fmt.Sprintf("worker %d processed task %d ", id, i)
		time.Sleep(time.Second)
	}
}

func main() {
	results := make(chan string, 10)

	//start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i, results)
	}

	//collect results
	for j := 1; j <= 20; j++ {
		fmt.Println(<-results)
	}

	fmt.Println("all tasks completed")

}
