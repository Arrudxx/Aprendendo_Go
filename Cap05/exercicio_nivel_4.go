package main

import "fmt"

func main() {

	x := 100
	fmt.Printf("%d, %b, %#x\n", x, x, x)

	y := x << 2
	fmt.Printf("%d, %b, %#x", y, y, y)

}
