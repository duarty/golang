package main

import "fmt"

func main(firstname = string) {
	var firstname string = "José"
	fmt.Println(firstname)

	lastname1 := "Duarte"
	fmt.Println(lastname1)

	lastname2, lastname3 := "Maduro", "Neto"
	fmt.Println(lastname2, lastname3)

	var (
		job   string = "Developer:"
		stack string = "JavaScript | Golang | Python"
	)
	println(job, stack)

	firstname = "first name mod"
	println(firstname)
}
