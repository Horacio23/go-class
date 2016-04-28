package main

import (
	"fmt"
)

func main() {
	//there are three ways to declare a variable
	var name string = "Horacio" // this is the complete way to declare a variable
	var name1 = "Alex"          // this is an inferred type of declaring a variable
	name3 := "Sandra"           // this is the other inferred type of declaring and issuing a value on the same line, it is the preffered method

	fmt.Println(name, name1, name3)
}
