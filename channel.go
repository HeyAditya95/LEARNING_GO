package main

import "fmt"

func main() {
	order := make(chan string) //here we creating a channel
	//chef preparing food
	go func() {
		order <- "pizza is ready" //sending data into channel
	}()

	//here we receiving data from the channel
	readyorder := <-order
	fmt.Println(readyorder)
}
