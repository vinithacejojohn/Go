package main

import "fmt"

func main() {
	// var card string = "Ace of Cards"
	card := newCard()
	fmt.Println(card)

}

func newCard() string {

	return "Ace of Cards"
}
