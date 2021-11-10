package main

import "fmt"

var number int

func init() {
	fmt.Println("funcão init")
	number = 10
}

func main() {
	fmt.Println("funcão main")
	fmt.Println(number)
}
