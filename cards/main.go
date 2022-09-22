package main

func main()  {
	cards := newDeck()

	// hand, remaniningCards := deal(cards, 5)
	// hand.print()
	// remaniningCards.print()
	// teste.TestePacote()
	cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
}
