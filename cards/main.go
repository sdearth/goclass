package main

func main() {
	//create a slice of type string (slice is an array with variable length)
	cards := newDeck()
	cards.saveToFile("myCards")

}
