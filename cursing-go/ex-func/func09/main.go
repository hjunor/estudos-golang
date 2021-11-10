package main

import "fmt"

func iverte(number int) int {
	return number * -1
}
func ivertePoiter(number *int) {
	*number = *number * -1
}
func main() {
	fmt.Println("funções com ponteiros")
	number := 12
	fmt.Println(number)
	fmt.Println(iverte(number))
	fmt.Println(number)
	newNumber := 10
	fmt.Println(newNumber)
	ivertePoiter(&newNumber)
	fmt.Println(newNumber)

}
