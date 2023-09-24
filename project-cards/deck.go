package main

import "fmt"

// A new "type" that "extends" the behavior of a slice of strings
type deck []string

// Generate new deck of cards
func newDeck() deck {
	fmt.Println("Creating new deck...")
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// A receiver function that allows all values of type deck to have access to printDeck()
func (d deck) printDeck() {
	fmt.Println("Printing deck...")
	for _, card := range d {
		fmt.Println(card)
	}
	fmt.Println("*********")
}

// A receiver function that allows all values of type deck to have access to print()
func (d deck) print() {
	fmt.Println("Calling print()...")
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// A function that deals a given number of cards from the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}
