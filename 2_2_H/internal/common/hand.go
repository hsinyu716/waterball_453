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
	h.cards = append(h.cards, card)
}

func (h *Hand) Show(index int) *Card {
	card := h.cards[index]
	h.cards = append(h.cards[:index], h.cards[index+1:]...)
	return card
}
