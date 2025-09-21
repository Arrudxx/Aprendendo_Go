package main

import "fmt"

func main() {
	x := 10
	//Operador de negação em go: !
	if !(x < 100) {
		fmt.Println("Hello World")
	}
}
