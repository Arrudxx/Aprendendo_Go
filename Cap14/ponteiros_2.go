package main

import "fmt"

func main() {
	x := 0

	EstaRecebeoValor(x)
	EstaRecebeUmPonteiro(&x)

}

func EstaRecebeoValor(x int) {
	x++
	fmt.Println("Na função:", x)
}

func EstaRecebeUmPonteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
