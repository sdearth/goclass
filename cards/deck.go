package main

import "fmt"

//Create a new type, deck, which is a slice of strings
//not an object, it just adds features to slice
type deck []string

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
