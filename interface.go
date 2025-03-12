package main

import "C"
import "fmt"

type Animal interface {
	speak() string
}
type Cat struct {
	name string
}
type dog struct {
	name string
}

func (c Cat) speak() string {
	return fmt.Sprintf("the name of the cat is %s and the sound is meow", c.name)
}
func (d dog) speak() string {
	return fmt.Sprintf("the name of the cat is %s and the sound is woof", d.name)
}
func main() {
	var a Animal
	cat1 := Cat{name: "tom"}
	dog1 := dog{name: "bob"}
	a = cat1
	a = dog1
	fmt.Println(a.speak())
	fmt.Println(a.speak())
}

//type shape interface {
//	area() float64
//}
//type circle struct {
//	radius float64
//}
//
//func (c circle) area() float64 {
//	return c.radius * c.radius * 3.14
//}
//
//func main() {
//	c1 := circle{10}
//	var s shape = c1
//	fmt.Println("the area of the circle is : ", s.area())
//}
