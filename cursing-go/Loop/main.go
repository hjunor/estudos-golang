package main

import (
	"fmt"
	"time"
)

func main() {

	var i int = 0
	for i < 10 {
		time.Sleep(1 * time.Second)
		println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		time.Sleep(1 * time.Second)
		println(j)
	}

	nomes := [3]string{"João", "Maria", "José"}

	for _, nome := range nomes {
		println(nome)
	}

	for i, char := range "Go lang" {
		fmt.Println(i, string(char))
	}

	user := map[string]string{
		"name":  "João",
		"email": "mail@mail.com",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

}
