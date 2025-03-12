package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	j := 1 //for Loop Without Initialization
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	//for { //for infinite loop
	//	fmt.Println("infinite loop")
	//}

	k := 1
	for k <= 5 {
		if k == 3 {
			k++
			continue
		}
		fmt.Println(k)
		k++
	}

	//if else
	var marks int = 24
	if marks >= 90 {
		fmt.Println("grade A")
	} else if marks >= 75 {
		fmt.Println("grade B")
	} else if marks >= 65 {
		fmt.Println("grade c")
	} else if marks < 33 {
		fmt.Println("tu fail hai bhen k lode")
	} else {
		fmt.Println("marks enter karo babuuuu")
	}

	//switch case
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("start of the week")
	case "tuesday":
		fmt.Println("2nd day of the week")
	case "wednesday":
		fmt.Println("3rd day of the week")
	case "thursday":
		fmt.Println("4th day of the week")
	case "friday":
		fmt.Println("frinday ayyy ")
	default:
		fmt.Println("weekend")
	}

	number := 10
	if number == 10 {
		fmt.Println("positive")
	} else if number == -10 {
		fmt.Println("negative")
	} else {
		fmt.Println("0")
	}

}
