package main

import "fmt"

type scopetNumber int

var y int
var x scopetNumber

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	y = int(x)
	fmt.Println(y)

}
