// 👉 “Fan-Out” means a single producer (main goroutine) sends tasks to multiple worker goroutines.
// 👉 Multiple goroutines process tasks concurrently, improving efficiency.

// 🔵 Real-Life Analogy for Fan-Out:

// Imagine a restaurant kitchen 🍽️.
//   - One chef (main goroutine) receives multiple orders.
//   - He assigns different cooks (worker goroutines) to prepare food simultaneously.
//   - This way, food is ready faster than if the chef cooked everything alone.
package main

import (
	"fmt"
	"time"
)

// worker function that processes tasks
func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d is performing the task %d\n", id, job)
		time.Sleep(time.Second) //stimualing processign time
	}
}

func main() {
	jobs := make(chan int, 5) //job queue
	NumWorkers := 3           //number of workers

	// start multiple goroutines
	for i := 1; i <= NumWorkers; i++ {
		go worker(i, jobs)
	}

	// send jobs to workers
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)

	time.Sleep(3 * time.Second) //wait for the workers to complete

}
