package main

import "fmt"

//program has a new type called bot. if you are a type with a function
//called getGreeting, you are also of type bot
//Note: this is implicit - you don't have to explicitly link types. As
// long as the type implements the interface methods, it's automatic.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//VERY CUSTOM LOGIC FOR GENERATING GREETING
	return "Hi there"
}

//can omit receiver name if we're not using it
func (spanishBot) getGreeting() string {
	//VERY CUSTOM LOGIC FOR GENERATING GREETING
	return "Hola"
}
