package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	var option int
	fmt.Println("Please enter a mumber to search")
	fmt.Scanln(&option)
	fmt.Println(binarySearch(array, option))
}

func binarySearch(array []int, num int) (answer int) {
	mid := len(array) / 2
	if len(array) <= 1 {
		answer = array[0]
		fmt.Println("The answer is ", answer)
		return
	}
	if array[mid] > num {
		fmt.Println("Array when less", array)
		answer = binarySearch(array[:mid], num)
	} else {
		fmt.Println("Array when more", array)
		answer = binarySearch(array[mid:], num)
	}

	return
}
