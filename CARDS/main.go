package main

func main() {
	cards := newDeckFromFile("my_card")
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_card")
	// hand, rest := cards.deal(4)
	// // fmt.Println(cards)
	// hand.print()
	// rest.print()
}
