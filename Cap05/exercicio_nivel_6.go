package main

import "fmt"

const (
	_ = 2025 + iota
	b
	c
	d
	e
)

func main() {

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
