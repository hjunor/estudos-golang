package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controles")

	number := 10

	if number > 10 {
		fmt.Println("O número é maior que 10")
	} else if number == 10 {
		fmt.Println("O número é igual a 10")
	} else {
		fmt.Println("O número é menor que 10")
	}
	if outherNumber := number + 1; outherNumber > 10 {
		fmt.Println("O número é maior que 10")
	} else if outherNumber == 10 {
		fmt.Println("O número é igual a 10")
	} else {
		fmt.Println("O número é menor que 10")
	}
}
