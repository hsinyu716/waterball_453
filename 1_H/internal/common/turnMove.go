package common

type TurnMove struct {
	player        PlayerService
	exchangeHands *ExchangeHands
	showCard      *Card
}

type TurnMoveService interface {
	GetExchangeHands() *ExchangeHands
	GetPlayer() PlayerService
	SetShowCard(card *Card)
	GetShowCard() *Card
}

func NewTurnMove(player PlayerService, exchangeHands *ExchangeHands, showCard *Card) *TurnMove {
	return &TurnMove{
		player:        player,
		exchangeHands: exchangeHands,
		showCard:      showCard,
	}
}

func (t *TurnMove) GetExchangeHands() *ExchangeHands {
	return t.exchangeHands
}

func (t *TurnMove) GetPlayer() PlayerService {
	return t.player
}

func (t *TurnMove) SetShowCard(card *Card) {
	t.showCard = card
}

func (t *TurnMove) GetShowCard() *Card {
	return t.showCard
}
