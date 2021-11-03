package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {

	fmt.Println("Binary \t \t \t decimal")
	fmt.Printf("%b \t \t %d \n", KB, KB)
	fmt.Printf("%b \t \t %d \n", MB, MB)
	fmt.Printf("%b \t \t %d \n", GB, GB)
	fmt.Printf("%b \t \t %d \n", TB, TB)

}
