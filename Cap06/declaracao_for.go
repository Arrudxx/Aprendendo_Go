package main

import "fmt"

func main() {

	// jeito de fazer um while em go
	// x := 0
	// for x < 10 {
	// 	fmt.Println("x é menor que deis")
	// 	x++
	// }

	x := 0
	for {
		if x < 10 {
			fmt.Println("x é menor que dez")
			x++
		} else {
			fmt.Println("x é maior")
			break
		}
	}

}
