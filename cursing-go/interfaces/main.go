package main

import "fmt"

type Employee interface {
	salary(value int) int
	lastName(value string)
}

type Age int

func (a Age) lastName(value string) {
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
