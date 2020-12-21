package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

func main() {
	fmt.Println("Welcome!")

	book := Book{"Oka", "shahin"}
	book2 := Book{Title : "pera nai chill", Author : "prangon"}

	fmt.Printf("%+v\n", book)
	fmt.Printf("%v\n", book)
	fmt.Println(book)

	fmt.Println(book2)


	byteArray, err := json.Marshal(book)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
	fmt.Println(byteArray)
}