package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}
type portugueseBot string

func main() {
	var pb portugueseBot
	eb := englishBot{}
	sb := spanishBot{}
	pb = "Alou"

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(pb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (portugueseBot) getGreeting() string {
	return "Alou!"
}

func (englishBot) getGreeting() string { 
	// o valor no receiver pode ser removido se não usado, deiando apenas o tipo
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
