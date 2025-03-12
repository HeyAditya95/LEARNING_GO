package main

import "fmt"

//type person struct {
//	Name string
//	Age  int
//}
//
//func (p person) introduce() {
//	fmt.Println("hi my name is", p.Name, "and I am", p.Age, "years old.")
//}
//
//func main() {
//	p1 := new(person)
//	p1.Name = "Jack"
//	p1.Age = 23
//	p1.introduce()
//	//p1 := person{Name: "aditya", Age: 20}
//	//fmt.Println(p1.Name, p1.Age)
//	//fmt.Println(p1.Age)
//}

// task part
type Car struct {
	Brand string
	Year  int
}

func (c Car) Info() {
	fmt.Println("The Brand of the car is ", c.Brand, "and the model year of the car is ", c.Year)
}

func main() {
	Car1 := Car{Brand: "BMW", Year: 2010}
	Car2 := Car{Brand: "Audi", Year: 2004}
	Car1.Info()
	Car2.Info()

}
