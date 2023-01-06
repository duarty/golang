package main

import "fmt"

func sum(num1 int16, num2 int16) int16 {
	return num1 + num2
}

func calc(num1, num2 int16) (int16, int16) {
	return num1 + num2, num1 - num2
}

func main() {
	total := sum(120, 80)
	fmt.Println(total)

	sum, sub := calc(10, 20)
	fmt.Println(sum, sub)

	sum2, _ := calc(300, 100)
	fmt.Println(sum2)

	_, sub2 := calc(300, 100)
	fmt.Println(sub2)
}
