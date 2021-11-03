package main

import "fmt"

const (
	number int = 255
	name       = "Heberth"
)

func main() {
	if 255 == number {
		fmt.Println("The number is equal to 255")
	} else if number != 255 {
		fmt.Println("The number is not equal to 255")
		if number <= 255 && number >= 0 {
			fmt.Println("The number is less than 255")

		} else if number >= 255 && number <= 1000 {
			fmt.Println("The number is between 255 and 1000")
		} else {
			fmt.Println("The number is between than 1000")
		}
	}

}
