// package main
//
// import (
//
//	"fmt"
//	"time"
//
// )
//
// // Worker function: Processes tasks from the jobs channel and sends results to results channel
//
//	func worker2(id int, jobs <-chan int, results chan<- string) {
//		for job := range jobs {
//			fmt.Printf("Worker %d started task %d\n", id, job)
//			time.Sleep(2 * time.Second) // Simulating work
//			fmt.Printf("Worker %d finished task %d\n", id, job)
//			time.Sleep(2 * time.Second)
//			results <- fmt.Sprintf("Task %d processed by Worker %d", job, id)
//		}
//	}
//
//	func main() {
//		numWorkers := 3 // Number of workers
//		numJobs := 5    // Number of tasks
//
//		jobs := make(chan int, numJobs)       // Channel to hold jobs (task queue)
//		results := make(chan string, numJobs) // Channel to collect results
//
//		// Create worker goroutines
//		for i := 1; i <= numWorkers; i++ {
//			go worker2(i, jobs, results)
//		}
//
//		// Send jobs (tasks) to the jobs channel
//		for j := 1; j <= numJobs; j++ {
//			jobs <- j
//		}
//		close(jobs) // Close the jobs channel after sending all jobs
//
//		// Collect results
//		for r := 1; r <= numJobs; r++ {
//			fmt.Println(<-results)
//		}
//
//		fmt.Println("All tasks are completed!")
//	}
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//// Worker function: Processes tasks from the jobs channel and sends results to results channel
//func worker2(id int, jobs <-chan int, results chan<- string) {
//	fmt.Printf("Worker %d is ready and waiting for a job...\n", id) // Debug statement
//	for job := range jobs {
//		fmt.Printf("Worker %d received task %d\n", id, job) // Debug statement
//		time.Sleep(2 * time.Second)                         // Simulating work
//		fmt.Printf("Worker %d finished task %d\n", id, job) // Debug statement
//		results <- fmt.Sprintf("Task %d processed by Worker %d", job, id)
//	}
//	fmt.Printf("Worker %d has no more tasks, exiting...\n", id) // Debug statement
//}
//
//func main() {
//	numWorkers := 3 // Number of workers
//	numJobs := 5    // Number of tasks
//
//	jobs := make(chan int, numJobs)       // Channel to hold jobs (task queue)
//	results := make(chan string, numJobs) // Channel to collect results
//
//	// Create worker goroutines
//	for i := 1; i <= numWorkers; i++ {
//		go worker2(i, jobs, results)
//	}
//
//	// Send jobs (tasks) to the jobs channel
//	for j := 1; j <= numJobs; j++ {
//		fmt.Printf("Main: Sending task %d to workers\n", j) // Debug statement
//		jobs <- j                                           // Sending job `j` into the channel
//	}
//	close(jobs)                                                               // Close the jobs channel after sending all jobs
//	fmt.Println("Main: All tasks have been assigned, waiting for results...") // Debug statement
//
//	// Collect results
//	for r := 1; r <= numJobs; r++ {
//		result := <-results
//		fmt.Println("Main received:", result) // Debug statement
//	}
//
//	fmt.Println("All tasks are completed! Closing program.") // Debug statement
//}

// ------------------WORKPOOL WAITGROUP----------------------
package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function
func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark task as done when function ends

	fmt.Printf("Worker %d started working...\n", id)
	time.Sleep(2 * time.Second) // Simulating some work
	fmt.Printf("Worker %d finished work!\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Add a new worker
		go worker2(i, &wg)
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers finished, main function exiting!")
}
