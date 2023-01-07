package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int32
	array1[3] = 400
	fmt.Println(array1)

	var array2 [10]string
	fmt.Println(array2)

	array2[5] = "Duarte"
	fmt.Println(array2)

	array3 := [5]int8{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice1 := []int8{1, 2, 3}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	slice1 = append(slice1, 19)
	fmt.Println(slice1)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[2] = 100
	fmt.Println(slice2)

	slice3 := make([]float32, 5, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
