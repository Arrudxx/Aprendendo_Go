package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// array2 := append(array, 6)
	slice2 := append(slice, 6)

	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 13231
	fmt.Println(slice[3])
}
