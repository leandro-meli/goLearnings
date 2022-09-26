package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}

type square struct{
	sideLength float64
}

func main() {
	sq := square{
		sideLength: 4,
	}
	tr := triangle{
		base: 9,
		height: 12,
	}
	printArea(sq)
	printArea(tr)
}

func (t triangle) getArea() (float64) {
	return 0.5 * t.base * t.height
}

func (sq square) getArea() (float64) {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
