package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var number int = 10

	var pointer *int = &number

	fmt.Println(number)
	fmt.Println(*pointer)
	twoNumber := 20
	//trocar o valor da avariavel number pela variavel twoNumber sem passar valor direto
	*pointer = twoNumber

	fmt.Println(number)
	fmt.Println(*pointer)

}
