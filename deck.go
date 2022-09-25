package main

import "fmt"

// Create a new type of Deck
// which is a type of string

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}
