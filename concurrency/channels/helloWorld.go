package main

import "fmt"

func main() {
	ch := make(chan string)
	ch <- "Hello"
	fmt.Println(<-ch)
}
