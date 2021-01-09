package main

import "fmt"

// var decSize int = 20
//  you can declare a variable outside the function but not initilize

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	card := newCard()

	// card = "Five of Diamonds"
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
