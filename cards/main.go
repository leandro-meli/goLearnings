package main

import (
	"fmt"

	"github.com/leandro-meli/goLearnings/teste"
)

func main()  {
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i + 1, card)
	}
	fmt.Println(cards)
	teste.TestePacote()
}

func newCard() string {
	return "Five of Diamonds";
}
