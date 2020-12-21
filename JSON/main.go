package main

import (
	"fmt"
)

type Book struct {
	Title string
	Author string
}

func main() {
	fmt.Println("Welcome!")

	book := Book{"Oka", "shahin"}
	book2 := Book{Title : "pera nai chill", Author : "prangon"}

	fmt.Printf("%+v\n", book)
	fmt.Printf("%v\n", book)
	fmt.Println(book)

	fmt.Println(book2)

}