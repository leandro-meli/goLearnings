package main

import (
	"fmt"

	"github.com/leandro-meli/goLearnings/teste"
)

func main()  {
	card := newCard()
	fmt.Println(card)
	fmt.Println(getTitle())
	teste.TestePacote()
}

func newCard() string {
	return "Five of Diamonds";
}

func getTitle() string {
	return "Harry potter"
}