package main

import "fmt"

func main() {

	slice := []int{10, 12, 43, 2}

	total := somtatoria(slice...)
	fmt.Println(total)
}

func somtatoria(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
