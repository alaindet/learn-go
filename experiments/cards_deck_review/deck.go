package main

import (
	"errors"
	"fmt"
)

var suits DeckSuits = DeckSuits{
	"Hearts",
	"Diamonds",
	"Clubs",
	"Spades",
}

var values DeckValues = DeckValues{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

var (
	ErrInvalidQuantity   = errors.New("invalid quantity provided")
	ErrInsufficientCards = errors.New("insufficient number of cards left")
)

type Deck struct {
	Cards        []string
	ShuffledDeck []string
}

type DeckSuits [4]string

type DeckValues [13]string

func NewDeck() *Deck {

	cards := make([]string, 0, len(suits)*len(values))
	shuffledCards := make([]string, 0, len(suits)*len(values))

	for _, suite := range suits {
		for _, value := range values {
			card := fmt.Sprintf("%s of %s", value, suite)
			cards = append(cards, card)
		}
	}

	copy(shuffledCards, cards)

	return &Deck{Cards: cards, ShuffledDeck: cards}
}

func (d *Deck) Shuffle() {

	shuffled := make([]string, 0, len(d.Cards))

	for i := 0; i < len(d.Cards); i++ {
		randomIndex, _ := getRandomInt(0, len(d.Cards)-1)
		card, _ := at(d.Cards, randomIndex)
		shuffled = append(shuffled, card)
	}

	d.ShuffledDeck = shuffled
}

func (d *Deck) Draw(quantity int) ([]string, error) {
	err := d.CheckAvailableCards(quantity)

	if err != nil {
		return nil, err
	}

	cards := make([]string, 0, len(d.Cards))

	for i := 0; i < quantity; i++ {
		card := d.ShuffledDeck[i]
		cards = append(cards, card)
	}

	d.ShuffledDeck = d.ShuffledDeck[quantity:]

	return cards, nil
}

func (d *Deck) Burn(quantity int) error {
	err := d.CheckAvailableCards(quantity)

	if err != nil {
		return err
	}

	d.ShuffledDeck = d.ShuffledDeck[quantity:]
	return nil
}

func (d *Deck) CheckAvailableCards(quantity int) error {

	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	if len(d.ShuffledDeck) < quantity {
		return ErrInsufficientCards
	}

	return nil
}
