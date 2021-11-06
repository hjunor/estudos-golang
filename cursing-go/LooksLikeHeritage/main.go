package main

import "fmt"

type Peoope struct {
	name string
	age  uint8
}

type Student struct {
	Peoope
	cursos []string
	ref    int32
}

func main() {
	fmt.Println("Herança sqn")

	var s Student
	var p Peoope
	p = Peoope{"Joao", 20}
	s = Student{p, []string{"Go", "Python"}, 10}

	fmt.Println(s)
	fmt.Println("curso 1 " + s.cursos[0])
	fmt.Println([]string{s.cursos[0], s.cursos[1], "Java"})

}
