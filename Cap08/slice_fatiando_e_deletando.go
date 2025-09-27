package main

import "fmt"

func main() {

	sabores := []string{"peperonni", "mussarela", "quatro queijos", "abacaxi"}

	// Slice
	fatia := sabores[:]

	fmt.Println(fatia)

	sabores = append(sabores[0:1], sabores[3:]...)

	fmt.Println(sabores)

}
