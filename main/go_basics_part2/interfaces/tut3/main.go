package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

type square struct {
	sideLength float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

type triangle struct {
	base  float64
	hight float64
}

func (tri triangle) getArea() float64 {
	return 0.5 * (tri.base * tri.hight)
}

func main() {
	newSquare := square{
		sideLength: 10.1,
	}
	printArea(newSquare)

	newTriangle := triangle{
		hight: 2.0,
		base:  5.0,
	}
	printArea(newTriangle)
}
