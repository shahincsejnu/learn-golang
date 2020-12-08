package main

import "fmt"
import "math"

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2-x1
	b := y2-y1

	return math.Sqrt(a*a + b*b)
}

func rectangleArea(r Rectangle) float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}


func main(){
	fmt.Println("Let's start chapter 9")
	//
	//var rx1, ry1 float64 = 0, 0
	//var rx2, ry2 float64 = 10, 10
	//var cx, cy, cr float64 = 0, 0, 5
	//
	//fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	//fmt.Println(circleArea(cx, cy, cr))

	//var c Circle
	//c := new(Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}

	fmt.Println(circleArea(c))

	r := Rectangle{0, 0, 10, 10}

	fmt.Println(rectangleArea(r))

}