package main

func main()  {
	cards := newDeckFromFile("my_cards")

	// hand, remaniningCards := deal(cards, 5)
	// hand.print()
	// remaniningCards.print()
	// teste.TestePacote()
	// cards.saveToFile("my_cards")
	cards.print()
	// fmt.Println(cards.toString())
}
