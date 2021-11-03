package main

import "fmt"

func main() {
	s := `Hello,	World!`
	slice := []byte(s)
	fmt.Printf("%v\n%T", s, s)
	fmt.Printf("\n%v\n%T \n", slice, slice)

	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v)
	}
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x \n", s[i], s[i], s[i], s[i])
	}

}
