package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"http://golang.org",
		"https://amazon.com",
		"https://facebook.com",
		"https://stackoverflow.com",
	}

	//make a channel that can communicate strings
	//between routines
	c := make(chan string)

	for _, link := range links {
		//create a new go routine to run the checkLink
		//function on the current link
		go checkLink(link, c)
	}

	//wait for message on channel, then use as arg to Println
	// note that the channel receiving operation is blocking
	// using a range with a channel waits for the channel to
	// return a value and then assign to l (the for loop is an
	// infinite loop at this point)
	for l := range c {
		//call a function literal that sleeps for 5 seconds and then
		// checks the link. We need to pass the l value into the function
		// literal as an argument to avoid capturing the value
		// NB -> never reference the same variable in two different go routines
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//send message into channel
		c <- link
		return
	}
	fmt.Println(link, "is up")
	//send message into channel
	c <- link
}
