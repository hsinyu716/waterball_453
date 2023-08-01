package common

type TurnMove struct {
	player        IPlayer
	exchangeHands *ExchangeHands
	showCard      *Card
}

type ITurnMove interface {
	GetExchangeHands() *ExchangeHands
	GetPlayer() IPlayer
	SetShowCard(card *Card)
	GetShowCard() *Card
}

func NewTurnMove(player IPlayer, exchangeHands *ExchangeHands, showCard *Card) *TurnMove {
	return &TurnMove{
		player:        player,
		exchangeHands: exchangeHands,
		showCard:      showCard,
	}
}

func (t *TurnMove) GetExchangeHands() *ExchangeHands {
	return t.exchangeHands
}

func (t *TurnMove) GetPlayer() IPlayer {
	return t.player
}

func (t *TurnMove) SetShowCard(card *Card) {
	t.showCard = card
}

func (t *TurnMove) GetShowCard() *Card {
	return t.showCard
}
