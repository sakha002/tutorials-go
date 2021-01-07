package main

func main() {

	// cards := newDeck()
	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_deck_file")

	newCards := newDeckFromFile("./my_deck_file")

	newCards.print()

	newCards.shufle()

	newCards.print()

}
