package main

//func worker(ch chan string, wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Println("worker is waiting for a task...")
//	msg := <-ch
//	fmt.Println("worker recieved task :", msg)
//}
//func main() {
//	ch := make(chan string)
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go worker(ch, &wg) // start a goroutine
//
//	ch <- "process order #123" //send task
//	wg.Wait()                  //wait for worker to complete
//	fmt.Println("main finished sending task")
//}
//another way is to use a channel as a signal instead of waitgroup

import (
	"fmt"
)

func worker(ch chan string, done chan bool) {
	fmt.Println("worker is waiting for a task...")
	msg := <-ch
	fmt.Println("worker is done the task of", msg)
	done <- true
}

func main() {
	ch := make(chan string)
	done := make(chan bool)
	go worker(ch, done)
	ch <- "cooking the pizza"
	<-done
	fmt.Println("main finished sending task")
}
