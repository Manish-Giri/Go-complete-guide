package main

import "fmt"

func main() {
	// declare slice
	cards := []string{newCard(), "Five of hearts", "Six of spades"}

	// iterate over slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of diamonds"
}
