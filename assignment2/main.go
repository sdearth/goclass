package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	myTriangle := triangle{base: 10.0, height: 2.0}
	mySquare := square{sideLength: 5.0}

	printArea(myTriangle)
	printArea(mySquare)
}

func printArea(s shape) {
	fmt.Println("Area is:", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
