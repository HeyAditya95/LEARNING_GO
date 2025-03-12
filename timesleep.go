package main

import (
	"fmt"
	"time"
)

func main() {
	var something string = "something"
	fmt.Printf("the type of the something is: %T\nand its saying that soemthign is %s\n", something, something)
	time.Sleep(2000 * time.Millisecond)
	var numeric int = 42
	fmt.Printf("the type of numeric is %T and its value is %d\n", numeric, numeric)
	impliciti := "string"
	fmt.Printf("the type of impliciti is %T and its value is %s\n", impliciti, impliciti)
	time.Sleep(500 * time.Millisecond)
}
