package main

import (
	"fmt"
	"time"
)

func chef1(order1 chan string) {
	time.Sleep(2 * time.Second)
	order1 <- "pizza is ready"
}
func chef2(order2 chan string) {
	time.Sleep(1 * time.Second)
	order2 <- "burger is ready"
}
func chef3(order3 chan string) {
	time.Sleep(0 * time.Second)
	order3 <- "chips is ready"
}
func main() {
	order1 := make(chan string)
	order2 := make(chan string)
	order3 := make(chan string)

	go chef1(order1)
	go chef2(order2)
	go chef3(order3)

	select {
	case order := <-order1:
		fmt.Println(order)
	case order := <-order2:
		fmt.Println(order)
	case order := <-order3:
		fmt.Println(order)
		//case <-time.After(2 * time.Second):
		//	fmt.Println("timeout")
	default:
		time.Sleep(6 * time.Second)
		fmt.Println("no orders is reaady yet")

	}
}
