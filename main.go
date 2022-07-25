package main

import "fmt"

// creating interfact
type bot interface {
	getGreeting() string
}

type englishBot struct{} // declaring empty structs
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// Conceptually, the getGreetings() for each type will have very different logic from each other

func (eb englishBot) getGreeting() string {
	// Pretend custom logic for generating english greeting
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	// Pretend custom logic for generating spanish greeting
	return "Hola"
}

// On the other hand, our printGreeting() function will have very similar logic regardless of language
// Solution: use interfacing

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}