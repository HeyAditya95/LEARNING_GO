package main

import (
	"context"
	"fmt"
	"time"
)

func child(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //agr parent context cancel ho gya
			fmt.Println("child: parenet cancelled the context! exiting...")
			return
		default:
			fmt.Println("child : still running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background()) // parenet context banaya
	defer cancel()

	go child(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("Main cancelling the text now...")
	cancel()
	time.Sleep(2 * time.Second)
}
