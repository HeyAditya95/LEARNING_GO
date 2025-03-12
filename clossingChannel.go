package main

import (
	"fmt"
	"time"
)

// ----------------------CLOSING THE CHANNELS--------------
//
//	func main() {
//		orders := make(chan string, 4)
//		orders <- "order1"
//		orders <- "order2"
//		orders <- "order3"
//		orders <- "order4"
//
//		close(orders)
//		for order := range orders {
//			fmt.Printf("here the order %s\n", order)
//			time.Sleep(1 * time.Second)
//		}
//
//		time.Sleep(5 * time.Second)
//		fmt.Println("all order are done , we are closing the restaurant")
//	}
//
// ------------TIMERS IN CHANNELS----------------
//
//	func main() {
//		fmt.Println("chef starts preparing food")
//		timer := time.NewTimer(4 * time.Second)
//		<-timer.C
//		fmt.Println("chef finished preparing food")
//	}
//
// ----------------TICKER IN CHANNELS-----------------
func main() {
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("waiter checking tables")
			<-ticker.C
		}
		done <- true
	}()
	<-done
	ticker.Stop()
	fmt.Println("waiter shift is over")
}
