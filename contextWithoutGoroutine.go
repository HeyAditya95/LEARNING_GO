package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func makeRequest(ctx context.Context, url string) {
	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx) // Request me context attach kar diya

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Request successful with status:", resp.Status)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	makeRequest(ctx, "https://jsonplaceholder.typicode.com/todos/1")
}
