package common

type Hand struct {
	cards []*Card
	name  string
}

type HandService interface {
	AddCard(card *Card)
	Show() *Card
}

func NewHand(name string) *Hand {
	return &Hand{
		name: name,
	}
}

func (h *Hand) AddCard(card *Card) {
	if len(h.cards) > 13 {
		panic("over 13")
	}
	h.cards = append(h.cards, card)
}

func (h *Hand) Show() *Card {
	card := h.cards[0]
	h.cards = h.cards[1:]
	return card
}
