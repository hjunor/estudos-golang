package main

import "fmt"

func toRecover() {
	fmt.Println("tentativa de recuperar e execução")
	if r := recover(); r != nil {
		fmt.Println("Recuperado: ", r)
	}
}

func aprove(n1, n2 float64) bool {
	defer toRecover()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A media é exatamente 6")
}
func main() {
	fmt.Println("Funções panic e recover")
	fmt.Println("Aprovado? ", aprove(6, 6))
}
