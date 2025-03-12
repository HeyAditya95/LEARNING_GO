package main

import "fmt"

func main() {
	const pi float32 = 3.1415926
	const name1 string = "kartik"
	fmt.Println("the value of pi is :", pi)
	fmt.Println("the name is :", name1)

	const (
		name            = "aditya" //implicitily
		date    float32 = 2003     //emplicit typed
		favLang         = "golang"
		age             = 22
	)

	//var x int = date // ❌ ERROR: Can't assign float to int as it was explicitly typed earlier
	//fmt.Println(x)
	var y int = age
	fmt.Println(y)
	fmt.Println(name)
	fmt.Println(date)
	fmt.Println(favLang)

	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday)
}
