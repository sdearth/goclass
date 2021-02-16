package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

/*
Can also simply omit the contactInfo structure member name and it will
default to the type name
type person struct {
	firstName string
	lastName  string
	contactInfo
}
*/

func main() {
	//creation relying on order
	alex := person{"Alex", "Anderson", contactInfo{"alex@aol.com", 27703}}

	sherry := person{firstName: "Sherry",
		lastName: "Hedrick",
		contact: contactInfo{
			email:   "orangecyclelady@gmail.com",
			zipcode: 27703,
		}}

	//create a variable of type person with zero values. (string is empty,
	// int is 0, boolean is false...)
	var steve person
	steve.firstName = "Steve"
	steve.lastName = "Dearth"

	//ampersand operator gets address of the variable pointed at
	sherryPointer := &sherry
	fmt.Printf("%p\n", sherryPointer)
	sherryPointer.updateName("Dearth")

	//SHORTCUT: In go, if the receiver is a pointer to a type, then we can either
	// call the function with a pointer, or with a reference to the target type, and
	// go will "generate" the pointer for us
	alex.updateName("Arsy")

	alex.print()
	sherry.print()
	steve.print()
}

//Note that this doesn't work, because it passes by value
// func (p person) updateName(newLastName string) {
// 	p.firstName = newLastName
// }

//asterisk in this declaration indicates that this is a pointer to a type person
func (pointerToPerson *person) updateName(newLastName string) {
	//asterisk operator returns the value at the specified address
	(*pointerToPerson).lastName = newLastName
}

//You can create functions with structs as receivers
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
