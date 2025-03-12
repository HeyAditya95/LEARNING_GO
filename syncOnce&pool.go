//package main
//
//import (
//	"fmt"
//	"sync"
//)
////---------once----------------
//var once sync.Once
//
//func initialize() {
//	fmt.Println("initializing the resources")
//}
//
//func worker3(wg *sync.WaitGroup) {
//	defer wg.Done()
//	once.Do(initialize)
//	fmt.Println("worker doing its task...")
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	//start the multiple goroutine
//	for i := 1; i < 6; i++ {
//		wg.Add(1)
//		go worker3(&wg)
//	}
//	wg.Wait()
//	fmt.Println("all worker done")
//}

// -------------pool------------
package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new pool")
			return "New object"
		},
	}
	//put object on pool
	pool.Put("floating tube")
	fmt.Println("Getting", pool.Get())
	fmt.Println("Getting", pool.Get())

}
