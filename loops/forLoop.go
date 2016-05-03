package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Value for i is", i)
	}

	j := 0

	j++

	//you cant send noInc(j++)
	noInc(j)

}

func noInc(num int) {
	fmt.Println(num)
}
