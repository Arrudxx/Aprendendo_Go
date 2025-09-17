package main

import "fmt"

const (
	a = iota + 10
	_ = iota
	c = iota + 20
	_ = iota
	e = iota
)

func main() {
	fmt.Println(a, c, e)
}
