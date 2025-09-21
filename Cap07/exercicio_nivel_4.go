package main

import "fmt"

func main() {
	anonascimento := 2002
	anoatual := 2025
	for {
		if anonascimento > anoatual {
			break
		} else {
			fmt.Println(anonascimento)
			anonascimento++
		}
	}
}
