package main

import "fmt"

// A new "type" that "extends" the behavior of a slice of strings
type deck []string

// A receiver function that allows all values of type deck to have access to print()
func (d deck) print() {
	fmt.Println("Calling print()...")
	for i, card := range d {
		fmt.Println(i, card)
	}
}
