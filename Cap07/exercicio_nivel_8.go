package main

import "fmt"

func main() {

	x := 10

	switch {
	case x == 10:
		fmt.Println("Sim")
	case x != 10:
		fmt.Println("Não")
	}

}
