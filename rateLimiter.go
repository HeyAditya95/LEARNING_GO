////-------------SIMPLE RATE LIMITER-------------
//
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	limiter := time.Tick(2 * time.Second)
//
//	for i := 1; i <= 5; i++ {
//		<-limiter
//		fmt.Println("processed request", i, "at", time.Now().Format("2006-01-02 15:04:05"))
//	}
//}

// --------------BURSTY RATE LIMITING-----------
// what if we want to handle burst of request , i.e, 3 request immediately , then 1 request per second----------
package main

import (
	"fmt"
	"time"
)

func main() {
	burstlimiter := make(chan time.Time, 3)
	//fill the buffer with 3 tokens --manually
	for i := 0; i < 3; i++ {
		burstlimiter <- time.Now()
	}

	//create a ticker
	go func() {
		for t := range time.Tick(1 * time.Second) {
			burstlimiter <- t
		}
	}()

	//stimulating 6 request
	for i := 0; i < 6; i++ {
		<-burstlimiter
		fmt.Println("Processed request", i, "at", time.Now().Format("15:04:05"))
	}

}
