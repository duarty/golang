package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int16
}

type student struct {
	person
	class string
}

func main() {
	p1 := person{"José", "Duarte", 32}
	fmt.Println(p1)

	s1 := student{p1, "golang"}
	fmt.Println(s1)
	fmt.Println(s1.firstname)
}
