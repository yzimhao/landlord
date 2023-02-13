package landlord

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	fmt.Printf("新牌：%+v\n\n", deck)

	deck.Shuffle()
	fmt.Printf("洗牌后：%+v\n\n", deck)

	player1 := Hand(deck.Deal(17))
	player2 := Hand(deck.Deal(17))
	player3 := Hand(deck.Deal(17))

	lord := deck.Deal(3)
	fmt.Printf("player1: %v\n", player1.Sort())

	fmt.Printf("player2: %v\n", player2.Sort())
	fmt.Printf("player3: %v\n", player3.Sort())
	fmt.Printf("lord: %v\n", lord)
}
