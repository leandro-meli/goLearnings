package main

import "fmt"

func main() {
	four()
}

func one() {
	a := 42
	b := "james bond"
	c := true
	fmt.Println(a, b, c)
}

func three() {
	var x int = 32
	var y string = "james"
	var z bool = false

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}

func four() {
	type hotdog int
	var x hotdog
	var y int
	fmt.Printf("%T", x) // prints the type
	x = 42
	y = 33
	fmt.Println(x)
	x = hotdog(y)
	fmt.Println(x)
}
