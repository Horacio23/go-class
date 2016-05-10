package main

import "fmt"

func mycounter() func() {
	theCount := 0
	increment := func() {
		theCount++
		fmt.Println("The count is: ", theCount)
	}

	return increment
}

func main() {
	increment := mycounter()

	for i := 0; i < 5; i++ {
		increment()
	}
}
