package common

import "fmt"

type IPlayerUno interface {
	TakeTurnUno()
}

func (p *PlayerAdapter[T]) TakeTurnUno() {
	topCard := p.game.tableTopCard()
	fmt.Println(fmt.Sprintf("topCard %v", topCard.Translate()))
	for i, card := range p.GetHand().cards {
		if card.(*CardUno).CompareCard(topCard) {
			p.ShowCard(i)
			p.game.GetTrash().Push(card)
			return
		}
	}
	if p.game.GetDesk().Size() == 0 {
		for _, c := range p.game.GetTrash().cards {
			p.game.GetDesk().Push(c)
		}
		p.game.GetDesk().Shuffle()
		p.game.GetTrash().cards = nil
		p.game.GetTrash().Push(topCard)
	}
	card := p.game.GetDesk().DrawCard().(*CardUno)
	fmt.Println(fmt.Sprintf("抽卡 %v", card.Translate()))
	// 抽卡判斷可以出
	if card.CompareCard(topCard) {
		p.game.GetTrash().Push(card)
		return
	}
	p.AddHandCard(card)
	return
}
