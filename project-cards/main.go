package main

func main() {
	// declare slice
	cards := deck{newCard(), "Five of hearts", "Six of spades"}

	// add one more card
	cards = append(cards, "Two of clubs")

	// use type deck's function to print
	cards.print()
}

func newCard() string {
	return "Ace of diamonds"
}
