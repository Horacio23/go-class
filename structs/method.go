package main

import "fmt"

func main() {
	mp := messagePrinter{"foo"}
	mp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message)
}
