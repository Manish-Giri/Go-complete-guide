package main

import "fmt"

func main() {
	// generate deck of cards
	cards := newDeck()
	fmt.Println("Length of deck = ", len(cards))

	// use type deck's function to print
	cards.printDeck()
	fmt.Println("*********")
	// deal 5 cards and store the 2 hands separately
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Printing hand dealt...")
	hand.printDeck()
	fmt.Println("Length of hand = ", len(hand))
	fmt.Println("Printing remaining hand...")
	remainingDeck.printDeck()
	fmt.Println("Length of remaining hand = ", len(remainingDeck))

}
