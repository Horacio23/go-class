package main

import "fmt"

func main() {
	array := []string{"item1", "item2", "item3"}

	// the range returns 2 arguments, the first is the index and the second is the data from the array
	for i, data := range array {
		fmt.Println("index:", i, "value:", data)
	}

}
