package main

import "fmt"

func main() {
	// var card string = "Ace of Cards"
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {

	return "Ace of Cards"
}
