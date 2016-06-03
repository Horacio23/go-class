package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	c := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			c <- true
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-c
			}()
		}
	}

	fmt.Scanln()
}
