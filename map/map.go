package main

import "fmt"

func main() {
	user := map[string]string{
		"name":  "José",
		"alias": "Zeca",
	}
	fmt.Println(user["alias"])

	positions := map[int]string{
		1: "José",
		2: "Maximus",
		3: "Kermit",
	}
	fmt.Println(positions[2])

	user3 := map[string]map[string]string{
		"Personal Info": {
			"firstname": "José Duarte",
			"lastname":  "Neto",
		},
		"address": {
			"Rua":    "Lobo Dalmada",
			"Bairro": "SF",
		},
	}
	fmt.Println(user3)
	//delete a key
	delete(user3, "address")
	fmt.Println(user3)
}
