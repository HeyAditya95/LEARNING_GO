package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("task %d completed\n", id)
}
func main() {
	runtime.GOMAXPROCS(2) //use 2 cpu cores enabled parallelism
	var wg sync.WaitGroup
	wg.Add(3)

	//giving task
	for i := 1; i <= 3; i++ {
		go task(i, &wg)
	}
	wg.Wait()
	fmt.Println("all tasks completed successfully")

}
