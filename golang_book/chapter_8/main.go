package main

import "fmt"

func zero(ptr *int){
	*ptr = 0
}

func swap(ptr1 *int, ptr2 *int) {
	*ptr1, *ptr2 = *ptr2, *ptr1
}

func main(){
	fmt.Println("This chapter is about Pointers")

	x := 5

	zero(&x)

	fmt.Println(x)

	xptr := new(int)

	*xptr = 5

	fmt.Println(*xptr)

	xx := 1
	yy := 2

	swap(&xx, &yy)

	fmt.Println(xx, yy)
}