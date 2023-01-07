package main

import "fmt"

func main() {
	var msg string = "TNC"
	f := func(msg string) string {
		return msg
	}(msg)

	var msg2 string = "FDP"

	func(msg2 string) {
		fmt.Println(msg2)
	}(msg2)

	fmt.Println(f)
}
