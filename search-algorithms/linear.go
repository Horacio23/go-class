package main

import (
	"fmt"
)

func main() {
	array := []int{2, 5, 3, 2, 1, 5, 32}
	var option int
	fmt.Println("Please enter a mumber to search")
	fmt.Scanln(&option)
	linearSearch(array, option)
}

func linearSearch(array []int, num int) {
	for _, v := range array {
		if v == num {
			fmt.Println("We found the number ", num)
			return
		}
	}
	fmt.Println("Number is not in the array")

}
