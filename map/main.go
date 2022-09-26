package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#0000FF",
	}

	colors["white"] = "#ffffff" // dot notation não acessa chave, apenas square brackets

	// delete(colors, "white") // delete
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c { // color = key, hex = value
		fmt.Println("Hex code for", color, "is", hex)
	}
}
