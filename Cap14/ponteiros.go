package main

import "fmt"

func main() {
	x := 0

	y := &x
	// para usar o endereço da variavel use &
	// para ver oq tem dentro do endereço de memoria use *
	// fmt.Println(*y)

	*y++
	fmt.Println(*y)
	fmt.Printf("%T, %T\n", x, y)
	fmt.Print(x, y)
}
