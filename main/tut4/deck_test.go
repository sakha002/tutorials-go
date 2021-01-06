package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected last card to be Four of Diamonds but got %v", d[len(d)-1])
	}
}
