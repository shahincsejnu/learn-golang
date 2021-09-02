package main

import "fmt"

func average(s []float64) float64 {
	total := 0.0

	for _, value := range s {
		total += value
	}

	return total / float64(len(s))
}

func f() (r int) {
	r = 1
	return
}

func f2() (int, int) {
	return 5, 6
}

func add(nums ...int) int {
	total := 0

	for _, value := range nums {
		total += value
	}

	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func FibGen() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f2, f1 = (f1 + f2), f2

		return f1
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	xs := []float64{1, 2, 3, 4, 5}

	fmt.Println(average(xs))

	fmt.Println(f())

	//x, y := f2()

	fmt.Println(f2())

	fmt.Println(add(1, 2, 3, 4))

	x := []int{1, 1, 1, 1}

	fmt.Println(add(x...))

	add2 := func(x int, y int) int {
		return x + y
	}

	fmt.Println(add2(1, 2))

	val := 0
	increment := func() int {
		val++
		return val
	}

	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()

	fmt.Println(nextEven)

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fibo := FibGen()

	fmt.Println("Let's print fibo series")

	for i := 0; i < 10; i++ {
		fmt.Println(fibo())
	}

	fmt.Println(factorial(5))

	//defer func(){
	//	str := recover()
	//	fmt.Println(str)
	//}()
	//panic("got panicked")
	defer second()

	first()

	str := "help"

	func(str string) {
		fmt.Println(str)
	}(str)
}
