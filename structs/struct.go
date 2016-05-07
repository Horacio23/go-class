package main

import "fmt"

func main() {
	// these use local execution stack
	foo := myStruct{}
	foo.myField = "bar"

	bar := myStruct{"sushi"}

	fmt.Println(foo.myField)
	fmt.Println(bar.myField)

	// creates a pointer, but GO manages it when accessing myField so no need to dereference it
	// to use the heap stack use the new() func
	myFoo := new(myStruct)
	myFoo.myField = "bar"
	fmt.Println(myFoo)

	myOtherFoo := &myStruct{"bar"}
	fmt.Println(myOtherFoo)
}

type myStruct struct {
	myField string
}
