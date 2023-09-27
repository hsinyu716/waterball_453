package common

import "cosmos.cards.showdown/internal/common/card"

type Hand struct {
	Cards []card.Card
	name  string
}

type HandService interface {
	AddCard(card card.Card)
	Show() card.Card
}

func NewHand(name string) *Hand {
	return &Hand{
		name: name,
	}
}

func (h *Hand) AddCard(card card.Card) {
	h.Cards = append(h.Cards, card)
}

func (h *Hand) Show(index int) card.Card {
	card0 := h.Cards[index]
	h.Cards = append(h.Cards[:index], h.Cards[index+1:]...)
	return card0
}
