package main

import (
	"fmt"
	"runtime"
)

var x = 10
var y = 10.0

func main() {
	a := "e"
	b := "é"
	c := "é"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n", d, e, f)

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

	fmt.Println(runtime.GOOS)   // Para ver o sistema operacional
	fmt.Println(runtime.GOARCH) // Para ver a arquitetura

	// i := uint16(65536) //OverFlow
	i := uint16(65535)
	fmt.Println(i + 2)
}
