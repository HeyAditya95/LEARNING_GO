//package main
//
//import (
//	"fmt"
//)
//
//func chef(ch chan<- string) {
//	fmt.Println("chef: preparing pizza...")
//	ch <- "pizza is ready"
//}
//func waiter(ch <-chan string) {
//	order := <-ch
//	fmt.Println("waiter : Delivering order -", order)
//}
//func main() {
//	ch := make(chan string)
//	go chef(ch)
//	go waiter(ch)
//	//time.Sleep(1 * time.Second)
//}

package main

import (
	"fmt"
	"time"
)

func ordertaker(orders chan<- string) {
	fmt.Println("customer : i want a burger")
	time.Sleep(2 * time.Second)
	orders <- "burger"
}

func chef(orders <-chan string, prepared chan<- string) {
	order := <-orders
	fmt.Println("chef is preparing", order)
	time.Sleep(2 * time.Second)
	prepared <- order + " is ready"
}

func waiter(prepared <-chan string) {
	order := <-prepared
	fmt.Println("waiter is serving", order)
}

func main() {
	orders := make(chan string)
	prepared := make(chan string)

	go ordertaker(orders)
	go chef(orders, prepared)
	go waiter(prepared)
	time.Sleep(5 * time.Second)
}

//package main
//
//import (
//	"fmt"
//)
//
//// OrderTaker function (only sends data)
//func orderTaker(orders chan<- string) {
//	fmt.Println("Customer: I want a burger.")
//	orders <- "Burger" // ✅ Send order to the channel
//}

// Chef function (receives and processes order, then sends update)
//func chef(orders <-chan string, prepared chan<- string) {
//	order := <-orders // ✅ Receive order
//	fmt.Println("Chef: Preparing", order)
//	time.Sleep(2 * time.Second)
//	prepared <- order + " is ready!" // ✅ Send update
//}
//
//// Waiter function (only receives and delivers order)
//func waiter(prepared <-chan string) {
//	order := <-prepared // ✅ Receive prepared order
//	fmt.Println("Waiter: Serving", order)
//}
//
//func main() {
//	orders := make(chan string)
//	prepared := make(chan string)
//
//	go orderTaker(orders)
//	go chef(orders, prepared)
//	go waiter(prepared)
//
//	time.Sleep(5 * time.Second) // Wait for all routines to finish
//}
