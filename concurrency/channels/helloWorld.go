package main

import "fmt"

func main() {
	ch := make(chan string, 1) //hast to be buffered or it would block when putting in hello
	ch <- "Hello"
	fmt.Println(<-ch)
}
