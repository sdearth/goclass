package main

import "fmt"

//Create a new type, deck, which is a slice of strings
//not an object, it just adds features to slice
type deck []string

func newDeck() deck {
	cards := deck{}
	
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardRanks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	
	for _, suit := range cardSuits {
		for _, rank := range cardRanks {
			cards = append(cards, rank + " of " + suit)
		}
	}
}

func deal(d deck, int handSize) (deck, deck) {
	//slice range syntax is slice[beginInclusive:endExclusive]
	//can omit either to assume beginning or end
	return d[:handSize], d[handSize:]
}

//d is a receiver. Any variable of type 'deck' now gets access to the print
//method. Receiver sets up methods on types we create.
//**every variable of type deck can call this function on itself
//** the actual copy of the deck we're working with is available within the
//** function as a variable called 'd'
func (d deck) print() {
	//range returns <index, value>. If we don't want to use the index,
	// ignore it with an underscore to avoid compilation errors
	for _, card := range d {
		fmt.Println(card)
	}
}
