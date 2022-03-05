package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func (d deck) deal(handSize int) (deck, deck) {
	start := d[:handSize]
	rest := d[handSize:]
	return start, rest
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
