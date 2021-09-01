package main

import(
	"fmt"
	"math"
)

type circle struct {
	radious float64
}

type rectangle struct {
	width float64
	length float64
}

func (r rectangle) area() float64 {
	return r.width * r.length
}

func (c circle) area() float64 {
	return c.radious * c.radious * math.Pi
}

func main() {
	c1 := &circle{10}
	c2 := c1
	c2.radious = 20

	r1 := rectangle{5, 6}
	r2 := r1
	r2.width = 10

	fmt.Println(c1.area())
	fmt.Println(c2.area())

	fmt.Println(r1.area())
	fmt.Println(r2.area())
}