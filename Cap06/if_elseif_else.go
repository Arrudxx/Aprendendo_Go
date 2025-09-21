package main

import "fmt"

func main() {
	if x := 10; x > 100 {
		fmt.Println("x é maior que cem")
	} else if x < 10 {
		fmt.Printf("x não é maior que cem")
	} else {
		fmt.Println("x não é menor que dez nem maior que cem")
	}
}
