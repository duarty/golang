package main

import "fmt"

func main() {
	var num1 int8 = 127 //int8 range -127 to 127
	fmt.Println(num1)

	var bytenumber byte = 255 // byte have 0 to 255 range
	fmt.Println(bytenumber)

	var num2 int16 = -32768 //int16 -32768 to 32768
	fmt.Println(num2)

	var num3 int32 = -2147483648 //int32 -2147483648 to 2147483648
	fmt.Println(num3)
	var runenumber rune = -2147483648 //same that int32 but with another name(alias)
	fmt.Println(runenumber)

	var num4 int64 = -9223372036854775808 //int64 -9223372036854775808 to 9223372036854775808
	fmt.Println(num4)

	var nonnegativenum uint = 2147483648 //no negative numbers
	fmt.Print(nonnegativenum)

	var float float32 = 39999999999999999999999999999999999999.10 //RTX 4090 price
	fmt.Print(float)
}
