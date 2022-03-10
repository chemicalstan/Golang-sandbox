package main

import "fmt"

type bot interface {
	greeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	print(eb)
	print(sb)
}

func print(b bot) {
	fmt.Println(b.greeting())
}

func (englishBot) greeting() string {
	return "Hello there!"
}

func (spanishBot) greeting() string {
	return "Hola"
}

func (spanishBot) sayGoal() string {
	return "Golaso"
}
