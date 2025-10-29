package main

import "fmt"

func main() {

	total, quantos := soma2(1, 2, 3, 231)
	fmt.Println(total, quantos)
}

// Func comum
// x int, y int OU x int, y string
func soma(x, y int) int {
	return x + y
}

// Func com multiplos retornos
func soma2(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}
