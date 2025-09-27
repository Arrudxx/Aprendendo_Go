package main

import "fmt"

func main() {
	// slice := []string{"Banana", "Maçã", "Uva"}

	// for indice, valor := range slice {
	// 	fmt.Println("No Indice", indice, "temos o valor: ", valor)
	// }

	// // Errado
	// // slice[4] = "Melancia"

	// // Certo
	// slice = append(slice, "Melancia")

	// for _, valor := range slice {
	// 	fmt.Println("temos o valor: ", valor)
	// }

	slice := []int{20, 21, 22, 23}

	total := 0

	for _, valor := range slice {
		total += valor // mesma coisa que total = total + valor
		fmt.Println("O total é:", total)
	}

}
