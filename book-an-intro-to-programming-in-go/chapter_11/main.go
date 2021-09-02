package main

import "fmt"
import "github.com/shahincsejnu/golang_codes/book-an-intro-to-programming-in-go/chapter_11/my_package"

//import m "github.com/shahincsejnu/golang_codes/book-an-intro-to-programming-in-go/chapter_11/my_package"

func main() {
	a := 10
	b := 20

	//fmt.Println(m.Add(a, b))
	fmt.Println(my_package.Add(a, b))
}
