package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

type IPlayerUno interface {
	TakeTurnUno()
}

func (p *PlayerAdapter) TakeTurnUno() {
	topCard := p.game.tableTopCard()
	fmt.Println(fmt.Sprintf("topCard %v", topCard.Translate()))
	for i, card0 := range p.GetHand().Cards {
		if card0.(*card.Uno).CompareCard(topCard) {
			p.ShowCard(i)
			p.game.GetTrash().Push(card0)
			return
		}
	}
	if p.game.GetDeck().Size() == 0 {
		fmt.Println("牌堆沒牌，由棄牌區重洗！！！")
		p.game.GetDeck().Cards = p.game.GetTrash().Cards
		p.game.GetTrash().Cards = nil
		p.game.GetDeck().Shuffle()
		p.game.GetTrash().Push(topCard)
	}
	card0 := p.game.GetDeck().DrawCard().(*card.Uno)
	fmt.Println(fmt.Sprintf("抽卡 %v", card0.Translate()))
	// 抽卡判斷可以出
	if card0.CompareCard(topCard) {
		p.game.GetTrash().Push(card0)
		return
	}
	p.AddHandCard(card0)
	return
}
