package landlord

// Player represents a player in the game
type Player struct {
	Name  string
	Hand  Hand
	Score int
}

// NewPlayer creates a new player with a given name and deck of cards
func NewPlayer(name string, deck Hand) Player {
	return Player{
		Name:  name,
		Hand:  deck,
		Score: 0,
	}
}

// PlayCard plays a card from the player's hand
func (p *Player) PlayCard(card int) Card {
	playedCard := p.Hand[card]
	p.Hand = append(p.Hand[:card], p.Hand[card+1:]...)
	p.Score += playedCard.Value + 3
	return playedCard
}
