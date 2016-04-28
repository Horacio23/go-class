package main

import "fmt"

func main() {
	switch name := "alex"; name {
	case "alex":
		fmt.Println("My name is Horacio")
		fallthrough
	case "delgado":
		fmt.Println("My last name is Delgado")
	default:
		fmt.Println("The princess is in another castle")

	}
}
