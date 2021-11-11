package main

import "fmt"

func closure() func() {
	text := "dentro da função closure"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	fmt.Println("funções closure")

	text := "dentro da função main"

	fmt.Println(text)

	newFunc := closure()

	newFunc()

}
