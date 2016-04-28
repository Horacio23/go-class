package main

import "fmt"

func main() {
	i := "30"
	j := "20"

	if i < j {
		fmt.Println("I is less")
	} else if i > j {
		fmt.Println("J is less")
	} else {
		fmt.Println("They are equals or things dont work")
	}
}
