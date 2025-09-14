package main

import "fmt"

var x int //Declaração

var a int
var b float64
var c string
var d bool

func main() {
	x = 10 // Inicialização
	x = 20 // Atribuição

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

}
