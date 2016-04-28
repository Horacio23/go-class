package main

import "fmt"

func main() {

	switch name := "horacio"; name {
	case "horacio":
		fmt.Println("Hello, this actually works")
	case "alex":
		fmt.Println("My name is alex")
	default:
		fmt.Println("You broke the matrix")

	}
}
