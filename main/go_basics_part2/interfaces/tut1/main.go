package main

import "fmt"

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting() string {
	// custom logic for generting english bot greeting
	return "hi there"

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
