package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func printText(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}
func main() {
	fmt.Println("Variaticas")
	res := sum(1, 33, 56, 34, 78, 77)
	fmt.Println(res)
	printText("numbers: ", 1, 33, 56, 34, 78, 77)
}
