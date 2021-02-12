package main

func main() {
	//create a slice of type string (slice is an array with variable length)
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
