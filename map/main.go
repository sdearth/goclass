package main

import "fmt"

func main() {
	//declares a map with keys of type string and values of type string
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#ffffff",
		"white": "#000000",
	}

	colors["white"] = "#000000"

	printMap(colors)
	//declare an empty map
	//var colors map[string]string

	//create using the make function. Creates a small empty map
	//colors := make(map[string]string)

	delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
