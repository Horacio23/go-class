package main

import "fmt"

func main() {
	channel := make(chan int, 3)

	go func() {
		list := []int{1, 2, 3}
		for _, v := range list {
			channel <- v
		}
		// This will tell anyone waiting on the channel that is has been closed.
		close(channel)
	}()

	// Range lists take everything inside channels and wait until they are closed
	// or new items come if the channel is empty
	for v := range channel {
		fmt.Println("out of the channel", v)
	}
}
