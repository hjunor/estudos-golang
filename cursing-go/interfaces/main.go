package main

import "fmt"
type person struct {
  firstName string
  lastname string
  age int
}
type Employee interface {
	salary(value int) int
	lastName(value string)
}


func (a Age) lastName() {
  
	fmt.Println(a)
	fmt.Println(value)

}

func (a Age) salary(value int) int {
	return value * 20
}

func main() {
	var peopleAge Employee

	peopleAge = Age(2)

	peopleAge.lastName("John")

	fmt.Println(peopleAge.salary(10))
}
