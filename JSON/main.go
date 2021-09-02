package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Engineer bool   `json:"is_engineer"`
}

func main() {
	fmt.Println("Welcome!")

	author := Author{Name: "shahin", Age: 23, Engineer: true}

	book := Book{"Oka", author}
	book2 := Book{Title: "pera nai chill", Author: author}

	fmt.Printf("%+v\n", book)
	fmt.Printf("%v\n", book)
	fmt.Println(book)

	fmt.Println(book2)

	byteArray, err := json.Marshal(book)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
	//fmt.Println(byteArray)

	byteArray2, err := json.MarshalIndent(book2, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray2))
	//fmt.Println(byteArray2)
}
