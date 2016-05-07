package main

import "fmt"

func main() {
	emp := enhancedMessagePrinter{}
	emp.message = "foo"
	emp.printMessage()

	// this way is not recommended since u have to know the implementation details
	// use a constructor function instead of this
	emp1 := enhancedMessagePrinter{messagePrinter{"foo"}}
	emp1.printMessage()

}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message)
}

type enhancedMessagePrinter struct {
	//by naming the another struct, we automatically get access to its methdos
	messagePrinter
}
