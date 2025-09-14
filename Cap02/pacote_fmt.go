package main

import "fmt"

func main() {
	// x := "Olá bom dia\ncomo vai?\t espero que bem." // string literals
	// y := `Olá bom dia\ncomo vai?\t espero que bem.` // raw string literals

	x := "Oi"
	y := "Bom dia"
	z := fmt.Sprint(x, " ", y) // Transforma em string sua variavel
	fmt.Print(z)
	// fmt.Println(y)

}
