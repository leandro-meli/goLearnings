package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
		for _, suit := range cardSuits {
			for _, value := range cardValues {
				cards = append(cards, value+" of "+suit)
			}
		}
		return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
