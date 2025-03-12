package main

import (
	"context"
	"fmt"
)

type ctxKey string

func processRequest(ctx context.Context) {
	//now i am retrieving the value stored in the context
	userId1 := ctx.Value(ctxKey("userId"))
	fmt.Println("Processing request for user id :", userId1)
}
func main() {
	//create a background context and store a value in it
	ctx := context.WithValue(context.Background(), ctxKey("userId"), 12435)

	// pass the context with the stored value
	processRequest(ctx)
}
