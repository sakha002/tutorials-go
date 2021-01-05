package main

import "fmt"

// creates a new type of deck which is (extends) a slice of strings
type deck []string

func newDeck() deck {
	cardSuites := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
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
