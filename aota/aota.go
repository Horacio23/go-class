package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

// iota gets reset per constant block, meaning it has a non global scope
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
