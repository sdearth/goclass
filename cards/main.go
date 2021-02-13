package main

func main() {
	//create a slice of type string (slice is an array with variable length)
	cards := newDeck()
	
	hand, remainingCards = deal(cards, 5)

	hand.print()
	remainingCards.print()
	
}
