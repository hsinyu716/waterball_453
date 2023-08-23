package common

import (
	card2 "cosmos.cards.showdown/internal/common/card"
	"fmt"
)

type IPlayerUno interface {
	TakeTurnUno()
}

func (p *PlayerAdapter) TakeTurnUno() {
	topCard := p.game.tableTopCard()
	fmt.Println(fmt.Sprintf("topCard %v", topCard.Translate()))
	for i, card := range p.GetHand().Cards {
		if card.(*card2.CardUno).CompareCard(topCard) {
			p.ShowCard(i)
			p.game.GetTrash().Push(card)
			return
		}
	}
	if p.game.GetDesk().Size() == 0 {
		for _, c := range p.game.GetTrash().Cards {
			p.game.GetDesk().Push(c)
		}
		p.game.GetDesk().Shuffle()
		p.game.GetTrash().Cards = nil
		p.game.GetTrash().Push(topCard)
	}
	card := p.game.GetDesk().DrawCard().(*card2.CardUno)
	fmt.Println(fmt.Sprintf("抽卡 %v", card.Translate()))
	// 抽卡判斷可以出
	if card.CompareCard(topCard) {
		p.game.GetTrash().Push(card)
		return
	}
	p.AddHandCard(card)
	return
}
