package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("[^0-9\\.\\-]")
	if err != nil {
		log.Fatal(err)
	}

	string := reg.ReplaceAllString("hello there -34.1235", "")
	fmt.Println(string)
}
