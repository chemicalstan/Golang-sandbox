package main

// Create a new type of deck
// Which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spade", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
