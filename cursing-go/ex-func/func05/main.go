package main

import "fmt"

func fibonacci(posicion int) int {
	if posicion <= 1 {
		return posicion
	}
	return fibonacci(posicion-1) + fibonacci(posicion-2)
}

func main() {
	fmt.Println("função recursivas")

	posi := uint(50)

	for i := 0; i < int(posi); i++ {
		fmt.Printf("%d : %d \n", i, fibonacci(i))
	}

}
