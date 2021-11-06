package main

import (
	"fmt"
)

type User struct {
	name  string
	age   uint8
	adres Adress
}
type Adress struct {
	city   string
	state  string
	number uint8
}

func main() {
	fmt.Println("Arquivo struct")

	var u User

	u.name = "João"
	u.age = 30

	fmt.Println(u)

	newAdress := Adress{"Rio de Janeiro", "RJ", 123}

	newUserTwo := User{"Pedro", 21, newAdress}
	fmt.Println(newUserTwo)

	newUserThree := User{name: "Maria"}
	fmt.Println(newUserThree)
}
