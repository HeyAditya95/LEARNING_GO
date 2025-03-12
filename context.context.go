//// ------------------TIMEOUT-----------------
//package main
//
//import (
//	"context"
//	"fmt"
//	"time"
//)
//
//func fetchdata(ctx context.Context) {
//	select {
//	case <-time.After(3 * time.Second): // simulating a long operation
//		fmt.Println("Data fetched successfully")
//	case <-ctx.Done():
//		fmt.Println("Context cancelled", ctx.Err())
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	fmt.Println("fetching data...")
//	fetchdata(ctx)
//	fmt.Println("process finished...")
//}

//package main
//
//import (
//	"context"
//	"fmt"
//	"time"
//)
//
//func fetchData(ctx context.Context) {
//	select {
//	case <-time.After(3 * time.Second): // Simulating a long operation
//		fmt.Println("Data fetched successfully!")
//	case <-ctx.Done(): // Context timeout/cancel triggered
//		fmt.Println("Operation canceled:", ctx.Err())
//	}
//}
//
//func main() {
//	// Create a context with a timeout of 2 seconds
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel() // Ensure cleanup of resources
//
//	fmt.Println("Fetching data...")
//	fetchData(ctx)
//	fmt.Println("Process finished.")
//}

// --------------WITHCANCEL----------------
//package main
//
//import (
//	"context"
//	"fmt"
//	"time"
//)
//
//func worker4(ctx context.Context, id int) {
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Printf("worker %d done the work\n", id)
//			return
//		default:
//			fmt.Printf("worker %d workingre\n", id)
//			time.Sleep(500 * time.Millisecond)
//		}
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//
//	//start goroutine
//	go worker4(ctx, 1)
//	go worker4(ctx, 2)
//
//	time.Sleep(2 * time.Second) //let the worker run for 2 seconds
//	fmt.Println("Main : Cancelling the context...")
//	cancel()
//
//	time.Sleep(1 * time.Second) //wait to see if workers stopped the work
//	fmt.Println("Main : All worker stopped...")
//
//}

// ----------------DEADLINE-------------------
package main

import (
	"context"
	"fmt"
	"time"
)

func worker4(ctx context.Context, id int) {
	for {
		select {
		case <-time.After(3 * time.Second):
			fmt.Printf("worker %d completed the work", id)
		case <-ctx.Done():
			fmt.Println("worker ", id, "Deadline reached :", ctx.Err())
			return //exit the loop when context is cancelled
		}
	}
}

func main() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	//start goroutineœ
	go worker4(ctx, 1)
	go worker4(ctx, 2)

	//by ⏳ keep main() alive to see goroutine output
	time.Sleep(4 * time.Second)

}
