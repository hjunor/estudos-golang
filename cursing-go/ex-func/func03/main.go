package main

import "fmt"

func main() {

	fmt.Println("função anonima")
	func(text string) {
		fmt.Println(text)

	}("Hellow World")

	result := func(text string) string {
		return text

	}("Hellow World")

	fmt.Println(result)
}
