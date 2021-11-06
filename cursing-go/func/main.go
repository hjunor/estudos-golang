package main

import "fmt"

func Sum(x, y int) int {
	return x + y
}

func CalcNumber(x int, y int) (int, int) {

	subtracao := x - y
	adicao := x + y

	return subtracao, adicao

}
func main() {
	isSum := Sum(10, 30)

	fmt.Println(isSum)

	var f = func(text string) string {

		fmt.Println(text)

		return text
	}

	result := f("whats that is grover?")

	fmt.Println(result)

	add, sub := CalcNumber(10, 20)

	fmt.Println(sub)
	fmt.Println(add)

}
