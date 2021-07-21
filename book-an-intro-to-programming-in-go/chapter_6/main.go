package main

import "fmt"

func main() {
	// arrays
	var x [5]float64

	x[4] = 100

	fmt.Println(x)

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / 5)
	fmt.Println(total / float64(len(x)))

	total = 0

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	// slices

	var sl []float64

	sl = []float64{1, 2, 3}

	fmt.Println(sl)

	sl2 := make([]float64, 5)

	sl2 = x[:]

	fmt.Println(sl2)

	arr := [5]float64{1, 2, 3, 4, 5}

	sl3 := arr[0:5]

	fmt.Println(sl3)

	sl4 := append(sl, 4, 5, 6, 7)

	fmt.Println(sl4)

	//var sl5 []float64

	sl5 := make([]float64, 3)

	fmt.Println(sl, sl2)

	copy(sl5, sl2)

	fmt.Println(sl5)

	// maps

	var mp map[string]int
	mp = make(map[string]int)

	mp["key"] = 10

	fmt.Println(mp["key"])

	val, ok := mp["key"]

	fmt.Println(val, ok)

	delete(mp, "key")

	fmt.Println(mp["key"])

	val2, ok2 := mp["key"]

	fmt.Println(val2, ok2)

	// problems

	xx := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	mn := 1000

	for i := 0; i < len(xx); i++ {
		if mn > xx[i] {
			mn = xx[i]
		}
	}

	fmt.Println(mn)
}
