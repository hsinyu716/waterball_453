package domain

import (
	"cosmos.big2/internal/common/poker"
	"fmt"
	"sort"
)

type Hand struct {
	cards []*poker.Card
	name  string
}

type HandService interface {
	AddCard(card *poker.Card)
	GetCards() []*poker.Card
	PlayCard(cards []poker.Card)
	Size() int
	CardList()
	HasClub() (int, bool)
}

func NewHand(name string) *Hand {
	return &Hand{
		name: name,
	}
}

func (h *Hand) GetCards() []*poker.Card {
	return h.cards
}

func (h *Hand) AddCard(card *poker.Card) {
	if len(h.cards) > 13 {
		panic("over 13")
	}
	h.cards = append(h.cards, card)
	sort.Slice(h.cards, func(i, j int) bool {
		// 測資到最後一步這邊才調整了花色順序
		if h.cards[i].GetRank() == h.cards[j].GetRank() {
			return h.cards[i].GetSuit() < h.cards[j].GetSuit()
		}
		return h.cards[i].GetRank() < h.cards[j].GetRank()
	})
}

func (h *Hand) PlayCard(cards []*poker.Card) {
	for _, c := range cards {
		for hi, hc := range h.cards {
			if hc == c {
				h.cards = append(h.cards[:hi], h.cards[hi+1:]...)
				break
			}
		}
	}
}

func (h *Hand) Size() int {
	return len(h.cards)
}

func (h *Hand) CardList() {
	cardIndex := ""
	cardText := ""
	for i, card := range h.cards {
		cardIndex = fmt.Sprintf("%s%d    ", cardIndex, i)
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[card.GetSuit()], poker.RankMap[card.GetRank()])
	}
	fmt.Println(cardIndex)
	fmt.Println(cardText)
}

func (h *Hand) HasClub() (index int, club bool) {
	index = -1
	club = false
	for i, card := range h.cards {
		if card.GetRank() == 0 && card.GetSuit() == 0 {
			index = i
			club = true
			break
		}
	}
	return
}
