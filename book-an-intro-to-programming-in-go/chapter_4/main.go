package main

import "fmt"

func main() {
	var x string = "Hello World"

	fmt.Println(x)

	var y string

	y = "shahin"

	fmt.Println(y)

	fmt.Println(x == y)

	f()

	const xx int = 10

	var (
		a = 5
		b = 6
		c = 10
	)

	a += 0
	b += 0
	c += 0

	var input int

	fmt.Scanf("%d", &input)

	fmt.Println(input)
}

func f() {
	fmt.Println("I am here")
}
