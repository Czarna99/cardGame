package main

import (
	deck "Pawelek242/cardGame/Deck"
)

func main() {
	cards := deck.NewDeck()
	cards.SaveToFile("MyCards")

}
