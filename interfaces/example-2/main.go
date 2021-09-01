package main

import(
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

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
	c1 := circle{10}

	r1 := rectangle{5, 6}


}