package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	///pointer is a memory reference
	var var3 int
	var pointer *int

	fmt.Println(var3, pointer)

	var3 = 100
	pointer = &var3

	fmt.Println(var3, pointer, *pointer)

	var3 = 150
	fmt.Println(var3, pointer, *pointer)
}
