package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	var option int
	fmt.Println("Please enter a mumber to search")
	fmt.Scanln(&option)
	binarySearch(array, option)
}

func binarySearch(array []int, num int) {
	mid := len(array) / 2
	if len(array) <= 1 {
		if array[0] == num {
			fmt.Println("We found the number", num)
		} else {
			fmt.Println("We did not find the number")
		}
		return
	}
	if array[mid] > num {
		fmt.Println("Array when less", array)
		binarySearch(array[:mid], num)
	} else {
		fmt.Println("Array when more", array)
		binarySearch(array[mid:], num)
	}

}
