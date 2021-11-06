package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array is Slices")
	//array
	var array [5]string

	array[0] = "Hello"

	array = [5]string{array[0], "World", "Go", "Coding", "Fun"}
	fmt.Println(array)

	array2 := [5]string{"Hello", "World", "Go", "Coding", "Fun"}
	fmt.Println(array2)

	array3 := [...]string{"Hello", "World", "Go", "Coding", "Fun"}
	fmt.Println(array3)

	//slice

	slice := []string{"Hello", "World", "Go", "Coding", "Fun"}

	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array))
	slice = append(slice, "New")
	fmt.Println(slice)
}
