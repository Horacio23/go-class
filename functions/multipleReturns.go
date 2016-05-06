package main

import "fmt"

func main() {
	// unnamed return args
	uRet1, uRet2 := unnamed()
	fmt.Println(uRet1, uRet2)

	// named parameters
	nRet1, nRet2 := named()
	fmt.Println(nRet1, nRet2)
}

func unnamed() (string, string) {
	greeting := "hello"
	name := "Stranger"

	return greeting, name
}

func named() (greeting string, name string) {
	// When using named return statements, only the return keyword needs to be used
	greeting = "hello"
	name = "Stranger"

	return
}
