package main

func main() {
	// var card string = "Ace of Cards"
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
