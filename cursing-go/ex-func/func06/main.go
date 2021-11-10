package main

import "fmt"

func primary() {
	fmt.Println("primary")
}
func secundary() {
	fmt.Println("secundary")
}

func main() {
	defer primary()
	//defer = adiar a função
	secundary()
}
