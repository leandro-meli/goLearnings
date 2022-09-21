package main

import (
	"fmt"

	"github.com/leandro-meli/goLearnings/teste"
)

func main()  {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
	fmt.Println(cards)
	teste.TestePacote()
}

func newCard() string {
	return "Five of Diamonds";
}
