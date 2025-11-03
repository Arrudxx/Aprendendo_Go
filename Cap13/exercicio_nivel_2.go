package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println((funcao1(s1...)))
	fmt.Println((funcao2(s2)))

}

func funcao1(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func funcao2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
