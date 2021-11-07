package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	user := map[string]string{"nome": "Cristian", "idade": "23"}

	fmt.Println(user)
	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"people": {
			"nome":  "Cristian",
			"idade": "23",
		},
	}
	user2["city"] = map[string]string{
		"nome": "São Paulo",
	}
	fmt.Println(user2)
	delete(user2, "city")
	fmt.Println(user2)
}
