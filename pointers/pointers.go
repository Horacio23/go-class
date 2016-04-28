package main

import (
	"fmt"
)

func main() {
	movie := "Titanic"

	fmt.Println(movie)

	increaseManliness(&movie)

	fmt.Println(movie)
}

func increaseManliness(movie *string) {
	*movie = "Deadpool"

	fmt.Println("trying to change movie:", *movie)

}
