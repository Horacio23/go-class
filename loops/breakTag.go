package main

import (
	"fmt"
)

func main() {
mytag:
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i:", i, "j:", j)

			if i == j && i != 0 {
				break mytag
			}
		}
		fmt.Println("First loop ends nad i is:", i)
	}
	fmt.Println("This is the end")
}
