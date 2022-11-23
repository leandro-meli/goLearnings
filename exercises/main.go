package main

import "fmt"

func main() {
	iotaLib()
}

func one() {
	x := 35
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}

func rawStringLiteral() {
	x := `son las strings literals
  pero a cada enter
  una lina va abajo`
	fmt.Println(x)
}

func iotaLib() {
	const (
		a = 2022 + iota
		_
		b
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
