package main

import "fmt"

type dog struct {
	name   string
	race   string
	weight int8
	age    int8
	color  string
	owner  owner
}

type owner struct {
	name    string
	address address
}

type address struct {
	rua    string
	numero int16
}

func main() {
	var maximus dog
	maximus.name = "maximus"
	maximus.color = "yellow"
	maximus.race = "Labrador"
	maximus.weight = 35
	maximus.age = 2
	maximus.owner.name = "josé"
	maximus.owner.address = address{"Rua Lobo Dalmada", 599}

	fmt.Println(maximus)

	/*
		 	viralata := dog{
				"Vira",
				"latao",
				10,
				2,
				"black",

			}
			fmt.Println(viralata)
	*/
}
