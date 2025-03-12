package main

import "fmt"

func main() {
	personage := make(map[string]int)

	personage["aditya"] = 22
	personage["adi"] = 18
	personage["ananya"] = 18

	fmt.Println("the age of ananya is ", personage["ananya"])

	delete(personage, "adi")
	fmt.Println(personage)

	age, exits := personage["aditya"]
	if exits {
		fmt.Println("the age is ", age)
	} else {
		fmt.Println("the age is not exists")
	}

}
