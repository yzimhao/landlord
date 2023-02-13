package landlord

import (
	"fmt"
)

const (
	Spade = iota
	Heart
	Club
	Diamond
)

var suits = []string{"♠️", "❤️", "♣️", "♦️"}
var values = []string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}

// Card represents a single playing card
type Card struct {
	Suit  int
	Value int
}

// String returns a string representation of a card
func (c Card) String() string {
	return fmt.Sprintf("%s%s", values[c.Value], suits[c.Suit])
}
