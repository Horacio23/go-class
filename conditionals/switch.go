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

	swtichTest()

}

func swtichTest() {
	i := 0

	switch {
	case i < 1:
		fmt.Println("Hey i printed one")
	case i > 2:
		fmt.Println("And i am 2")
	}
}
