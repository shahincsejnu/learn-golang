package main

import "fmt"

var n int = 10

func main() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 = ", 1.1+1)
	fmt.Println("Hello dunia\n")
	fmt.Println(`different
					then double quote`)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x string = "hello"

	fmt.Println(x)

	y := "world"

	fmt.Println(y)

	z := 10

	z += 0

	const xx string = "constant"

	var (
		a = 5
		b = 10
		c = 20
	)

	a += 0
	b += 0
	c += 0

	fmt.Println("Enter a number : ")
	var input int

	fmt.Scanf("%d", &input)

}
