package landlord

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Hand represents a collection of cards
type Hand []Card

// NewDeck creates a new deck of cards
func NewDeck() Hand {
	deck := make([]Card, 0, 54)
	for suit := 0; suit < 4; suit++ {
		for value := 0; value < 13; value++ {
			deck = append(deck, Card{suit, value})
		}
	}
	deck = append(deck, Card{4, 13}) //小鬼
	deck = append(deck, Card{4, 14}) //大鬼
	return deck
}

// String returns a string representation of a hand of cards
func (h Hand) String() string {
	s := make([]string, len(h))
	for i, c := range h {
		s[i] = c.String()
	}
	return fmt.Sprintf("[%s]", s)
}

// Less is used for sorting the deck of cards
func (h Hand) Less(i, j int) bool {
	if h[i].Value == h[j].Value {
		return h[i].Suit < h[j].Suit
	}
	return h[i].Value < h[j].Value
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) Len() int {
	return len(h)
}

// Score returns the score of a hand of cards
func (h Hand) Score() int {
	score := 0
	for _, c := range h {
		score += c.Value + 3
	}
	return score
}

//洗牌
func (h *Hand) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*h), func(i, j int) {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	})
}

//发牌
func (h *Hand) Deal(n int) []Card {
	hand := (*h)[:n]
	*h = (*h)[n:]
	return hand
}

//排序
func (h Hand) Sort() Hand {
	sort.Sort(h)
	return h
}
