package  main

import "fmt"

func main() {
	fmt.Println(Total(1, 2, 3))

	fmt.Println(Xor(1, 2, 3))
}

func Total(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}

	return total
}

func Xor(args ...int) int {
	ans := 0

	for _, value := range args {
		ans ^= value
	}

	return ans
}