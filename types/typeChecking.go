package main

import "fmt"

type Shape struct {
	name string
}

func main() {

	var circle interface{} = Shape{"Circle"}

	if _, ok := circle.(*Shape); !ok {

		fmt.Println("hey circle is a shape")
	} else {
		fmt.Println("not a shape")

	}
}
