package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7}
	y := []int{33, 44, 55}

	x = append(x, y...)
	fmt.Printf("%v", x[4])
	// {4, 5, 6, 7, {33, 44, 55}}
}
