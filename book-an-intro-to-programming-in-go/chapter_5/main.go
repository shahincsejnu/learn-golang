package main

import (
	"fmt"
)

func main() {
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i += 1
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)

		if j%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	ch := 'a'

	switch ch {
	case 'b':
		fmt.Println("This is a consonant")
	case 'a':
		fmt.Println("This is a vowel")
	default:
		fmt.Println("Error")
	}
}
