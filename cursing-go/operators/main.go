package main

import (
	"fmt"
	"log"
)

type num interface {
	numOperators(int, int, int, int, int) int
}

func operations(a, b int) (int, int, int, int, int) {

	if a < b {
		log.Println("error")
		return 0, 0, 0, 0, 0
	} else {
		return a + b, a - b, a * b, a / b, a % b

	}
}

func main() {
	//aritimetic operators + - / * %
	sum, sub, mult, div, mod := operations(2, 5)

	fmt.Println(sum, sub, mult, div, mod)

	//operation atributes

	var variable string = "that is number"
	fmt.Println(variable)
	//variable = "that is string"

	variableThooseDays := "that is string toose day"
	fmt.Println(variableThooseDays)
	//variableThooseDays = "that is string toose day"
}
