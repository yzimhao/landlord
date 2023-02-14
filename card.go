package landlord

import (
	"fmt"
)

type suit int

const (
	SuitSpade suit = iota
	SuitHeart
	SuitClub
	SuitDiamond
	SuitJoker

	Smalljoker = "小王"
	Bigjoker   = "大王"
)

var suits = []string{"♠️", "❤️", "♣️", "♦️", "Joker"}
var values = []string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2", Smalljoker, Bigjoker}

// Card represents a single playing card
type Card struct {
	Suit  int
	Value int
}

// String returns a string representation of a card
func (c Card) String() string {
	return fmt.Sprintf("%s%s", values[c.Value], suits[c.Suit])
}

func (c Card) Point() string {
	return values[c.Value]
}

func parseCard(suit suit, value string) Card {
	pos := -1
	for i, v := range values {
		if v == value {
			pos = i
			break
		}
	}
	return Card{
		Suit:  int(suit),
		Value: pos,
	}
}
