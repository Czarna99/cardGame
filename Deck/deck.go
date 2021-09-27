package deck

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Create a new type of deck as slice of strings

type Deck []string

var err error

func NewDeck() Deck {

	cards := Deck{}

	cardSuits := []string{"Pik", "Kier", "Karo", "Trefl"}
	cardValues := []string{"As", "Dwójka", "Trójka", "Czwórka", "Piątka", "Szóstka", "Siódemka", "Osemka", "Dziewiątka", "Dziesiatka", "Walet", "Dama", "Król"}

	for _, suit := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" "+suit)
		}
	}

	return cards
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], nil

}

func (d Deck) Print() {
	for _, card := range d {

		fmt.Println(card)
	}
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), ", ")

}

func (d Deck) SaveToFile(filename string) error {

	ioutil.WriteFile(filename, []byte(d.ToString()), 0666)

	if err != nil {
		return err
	}

	return nil

}

func NewDeckFromFile(filename string) (Deck, error) {
	cards, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	newCards := string(cards)

	cardSlice := strings.Split(newCards, ", ")

	if err != nil {
		return nil, err
	}
	d := Deck(cardSlice)

	return d, nil
}
