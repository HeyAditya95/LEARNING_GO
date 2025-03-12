package main

import (
	"fmt"
	"sync"
	"time"
)

//var counter int64
//
//func increment(wg *sync.WaitGroup) {
//	defer wg.Done()
//	for i := 0; i < 1000; i++ {
//		atomic.AddInt64(&counter, 1) //atomic counts
//	}
//}
//
//// with atomic count
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go increment(&wg)
//	go increment(&wg)
//
//	wg.Wait()
//	fmt.Println("final counter is", counter)
//
//}

//with race condition
//func main() {
//	go increment()
//	go increment()
//	time.Sleep(1 * time.Second)
//	fmt.Println("final Counter:", counter)
//}

//real world example - counting website visits in a high traffic web server

// --------------------------with mutex-----------------
var counter int
var mu sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}
func main() {
	go increment()
	go increment()
	time.Sleep(1 * time.Second)
	fmt.Println("FINAL Counter:", counter)
}
