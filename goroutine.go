package main

import "fmt"

// function with parameter
func greet(name string) {
	fmt.Println("Hello", name)
}

// function with returning value
func add(a int, b int) int {
	return a + b
}

// multiple returning value
func calculate(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product

}

func main() {
	greet("aditya")
	sum := add(1, 2)
	fmt.Println("sum is", sum)
	multiply := sum * 2
	fmt.Println("multiply is", multiply)

	fmt.Println("-------------------returning mutiple values-------------------")
	sum, product := calculate(5, 7)
	fmt.Println("sum is", sum)
	fmt.Println("product is", product)

}
