package main

import "fmt"

// creating interface
type bot interface {
	getGreeting() string // anyone with this function can now call printGreeting()

}

type englishBot struct{} // declaring empty structs
type spanishBot struct{} // these types will have different getGreeting() fn's 

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

// any type with getGreeting() can call this
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}