package main

import (
	"errors"
	"fmt"
)

type Deck []string

type DeckSuits [4]string

type DeckValues [13]string

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

func NewDeck() *Deck {

	cardsCount := len(suits) * len(values)
	orderedCards := make([]string, 0, cardsCount)
	shuffledCards := make([]string, 0, cardsCount)

	// Get all available cards
	for _, suite := range suits {
		for _, value := range values {
			card := fmt.Sprintf("%s of %s", value, suite)
			orderedCards = append(orderedCards, card)
		}
	}

	// Shuffle cards
	for i := 0; i < cardsCount; i++ {
		randomIndex, _ := getRandomInt(0, len(orderedCards)-1)
		card, _ := at(orderedCards, randomIndex)
		shuffledCards = append(shuffledCards, card)
		orderedCards, _ = removeAt(orderedCards, randomIndex)
	}

	deck := Deck(shuffledCards)
	return &deck
}

func (d *Deck) Draw(quantity int) ([]string, error) {
	err := d.CheckAvailableCards(quantity)

	if err != nil {
		return nil, err
	}

	cards := make(Deck, quantity)
	copy(cards, *d)
	*d = Deck(*d)[quantity:]

	return cards, nil
}

func (d *Deck) Burn(quantity int) error {
	err := d.CheckAvailableCards(quantity)

	if err != nil {
		return err
	}

	*d = []string(*d)[quantity:]

	return nil
}

func (d *Deck) CheckAvailableCards(quantity int) error {

	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	if len(*d) < quantity {
		return ErrInsufficientCards
	}

	return nil
}
