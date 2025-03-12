package main

import (
	"fmt"
	"sync"
	"time"
)

var data int
var rw sync.RWMutex

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.RLock()
	fmt.Printf("reader %d is reading : %d\n", id, data)
	rw.RUnlock()
}

func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.Lock()
	fmt.Printf("writer %d is writing : %d\n", id, data)
	data += 10
	time.Sleep(1 * time.Second)
	fmt.Printf("writer %d is writing : %d\n", id, data)
	rw.Unlock()
}

func main() {
	var wg sync.WaitGroup
	data = 100

	//start multiple agent reader
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}

	//start a writer
	wg.Add(1)
	go writer(1, &wg)

	//start more reader
	for i := 4; i <= 6; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}
	wg.Wait()
	fmt.Println("FINAL Counter:", data)
}
