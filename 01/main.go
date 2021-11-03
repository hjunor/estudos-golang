package main

// Language: go
// Path: 02/main.go

import "fmt"

var primaryNumber = []int{1, 2, 3, 4, 5}
var primaryNumberLine = 20
var secondNumber = 30

func main() {

	var result = multiplierNumber(primaryNumberLine, secondNumber)
	var addResult = addNumber(primaryNumberLine, secondNumber)
	var sub = subNumber(primaryNumberLine, secondNumber)
	fmt.Println((result))
	fmt.Println(sub)
	fmt.Println((addResult))

	var arrayLine = arrayAppend(primaryNumber, secondNumber)
	var result2 = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(arrayLine)
	fmt.Println(result2)
}

func multiplierNumber(primaryNumber int, secondNumber int) int {
	return primaryNumber * secondNumber
}

func addNumber(primaryNumber int, secondNumber int) int {
	return primaryNumber + secondNumber
}
func subNumber(primaryNumber int, secondNumber int) int {
	return primaryNumber - secondNumber
}

func arrayAppend(primaryNumber []int, secondNumber int) []int {
	return append(primaryNumber, secondNumber)
}
