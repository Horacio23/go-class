package main

import (
	"fmt"
	"sync"
)

func main() {
	// The waitgroup is used to allow the go routines to finish before main exits
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		//defer is executed only after a function is done
		defer waitGrp.Done()
		fmt.Println("Inside the internal function 1")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Inside the internal function 2")
	}()

	waitGrp.Wait()
}
