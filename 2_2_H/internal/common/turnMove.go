package common

import (
	"cosmos.cards.showdown/internal/common/card"
)

type TurnMove struct {
	player        IPlayer
	exchangeHands *ExchangeHands
	showCard      card.Card
}

type ITurnMove interface {
	GetExchangeHands() *ExchangeHands
	GetPlayer() IPlayer
	SetShowCard(card card.Card)
	GetShowCard() card.Card
}

func NewTurnMove(player IPlayer, exchangeHands *ExchangeHands, showCard card.Card) *TurnMove {
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

func (t *TurnMove) SetShowCard(card card.Card) {
	t.showCard = card
}

func (t *TurnMove) GetShowCard() card.Card {
	return t.showCard
}
