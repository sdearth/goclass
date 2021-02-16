package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type, deck, which is a slice of strings
//not an object, it just adds features to slice
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardRanks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, rank := range cardRanks {
			cards = append(cards, rank+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	//slice range syntax is slice[beginInclusive:endExclusive]
	//can omit either to assume beginning or end
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//use type coercion to convert a deck to a string slice, then
	//join the slice into a comma separated string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bytes), ",")
	return deck(s)
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

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
