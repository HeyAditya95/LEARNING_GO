package main

import (
	"fmt"
	"time"
)

type Request struct {
	amount int
	reply  chan int
}

func cashier(sales chan Request) {
	totalsales := 0 // this is our shared state

	for req := range sales {
		if req.amount == 0 {
			req.reply <- totalsales //send current sales if amount is 0
		} else {
			totalsales += req.amount //update sales
			req.reply <- totalsales  //send updates
		}
	}
}
func main() {
	sales := make(chan Request)
	go cashier(sales) //start cashier goroutine

	//simulate customer making purchase
	customer1 := make(chan int)
	sales <- Request{amount: 500, reply: customer1}
	fmt.Println("updated sales ", <-customer1)

	customer2 := make(chan int)
	sales <- Request{amount: 700, reply: customer2}
	fmt.Println("updated sales ", <-customer2)

	//check total sales
	customer3 := make(chan int)
	sales <- Request{amount: 0, reply: customer3}
	fmt.Println("total sales ", <-customer3)

	time.Sleep(1 * time.Second) //give some time before program exits
}
