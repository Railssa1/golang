package main

import (
	"fmt"
)

func main() {
	// Array
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slice

	slice := []int{6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice = append(slice, 11)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Array interno
	fmt.Println("-----------------")
	slice3 := make([]float32, 5, 10)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
}
