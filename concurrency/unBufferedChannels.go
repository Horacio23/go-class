package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 2
	}()

	// the routine will block here until there is something put into the channel
	num := <-channel
	fmt.Println(num)
}
