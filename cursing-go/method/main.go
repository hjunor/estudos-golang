package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	heigth    string
}

func (p Person) getName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) legalAge() bool {

	return p.age >= 18

}
func (p *Person) addAge() {
	p.age++
}

func main() {
	p1 := Person{"James", "Bond", 32, "1.75"}
	fmt.Println("Methodos")

	fmt.Println(p1.getName())
	fmt.Println(p1.legalAge())
	fmt.Println(p1.age)
	p1.addAge()
	fmt.Println(p1.age)

}
