package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type PlayerAdapter struct {
	name          string
	point         int
	showdown      *Showdown
	hand          *Hand
	exchangeHands *ExchangeHands
	uno           *Uno
}

type IPlayer interface {
	NameHimself(i int)
	SetShowdown(showdown_ *Showdown)
	SetHand(hand *Hand)
	AddHandCard(card *Card)
	GetHand() *Hand
	TakeTurn() *TurnMove
	MakeExchangeHandsDecision() *ExchangeHands
	HasUsedExchangeHands() bool
	GetExchangeHands() *ExchangeHands
	SetExchangeHands(exchangeHands *ExchangeHands)
	ShowCard(index int) *Card
	GainPoint()
	GetPoint() int
	GetName() string
	GetCardSize() int

	SetGame(uno interface{})
	TakeTurnUno()
}

func (p *PlayerAdapter) NameHimself(i int) {
}

func (p *PlayerAdapter) SetShowdown(showdown *Showdown) {
	p.showdown = showdown
}

func (p *PlayerAdapter) SetHand(hand *Hand) {
	p.hand = hand
}

func (p *PlayerAdapter) AddHandCard(card *Card) {
	p.hand.AddCard(card)
}

func (p *PlayerAdapter) GetHand() *Hand {
	return p.hand
}

func (p *PlayerAdapter) TakeTurn() *TurnMove {
	return nil
}

func (p *PlayerAdapter) MakeExchangeHandsDecision() *ExchangeHands {
	return nil
}

func (p *PlayerAdapter) HasUsedExchangeHands() bool {
	return p.exchangeHands != nil
}

func (p *PlayerAdapter) GetExchangeHands() *ExchangeHands {
	return p.exchangeHands
}

func (p *PlayerAdapter) SetExchangeHands(exchangeHands *ExchangeHands) {
	p.exchangeHands = exchangeHands
}

func (p *PlayerAdapter) ShowCard(index int) *Card {
	return p.GetHand().Show(index)
}

func (p *PlayerAdapter) GainPoint() {
	p.point++
}

func (p *PlayerAdapter) GetPoint() int {
	return p.point
}

func (p *PlayerAdapter) GetName() string {
	return p.name
}

func (p *PlayerAdapter) GetCardSize() int {
	return len(p.GetHand().cards)
}

func (p *PlayerAdapter) filterOtherPlayer() []IPlayer {
	var selectPlayers []IPlayer
	for _, player := range p.showdown.GetPlayers() {
		if p.name != player.GetName() {
			selectPlayers = append(selectPlayers, player)
		}
	}
	return selectPlayers
}

func (p *PlayerAdapter) selectExchangeHandsTarget(players []IPlayer) *ExchangeHands {
	printPlayerChoices(players)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	input = strings.TrimSuffix(input, "\n")
	targetIndex, err := strconv.Atoi(input)
	if targetIndex >= len(players) || targetIndex < 0 {
		return p.selectExchangeHandsTarget(players)
	}
	exchangeHands := NewExchangeHands(p, players[targetIndex])
	return exchangeHands
}

func printPlayerChoices(players []IPlayer) {
	str := "\n"
	for i, player := range players {
		str += fmt.Sprintf("(%d) %s \n", i, player.GetName())
	}
	fmt.Println(fmt.Sprintf("Select the target %s", str))
}

func (p *PlayerAdapter) SetGame(uno interface{}) {
	p.uno = uno.(*Uno)
}

func (p *PlayerAdapter) TakeTurnUno() {
	topCard := p.uno.tableTopCard()
	fmt.Println(fmt.Sprintf("topCard %v", topCard.translate()))
	for i, card := range p.GetHand().cards {
		if p.compareCard(topCard, card) {
			p.ShowCard(i)
			return
		}
	}
	if p.uno.desk.Size() == 0 {
		for _, c := range p.uno.pool.cards {
			p.uno.desk.Push(c)
		}
		p.uno.desk.Shuffle()
		p.uno.pool.cards = nil
		p.uno.pool.Push(topCard)
	}
	card := p.uno.desk.DrawCard()
	fmt.Println(fmt.Sprintf("抽卡 %v", card.translate()))
	// 抽卡判斷可以出
	if p.compareCard(topCard, card) {
		return
	}
	p.AddHandCard(card)
	return
}

func (p *PlayerAdapter) compareCard(topCard *Card, card *Card) bool {
	if topCard.color == card.color || topCard.number == card.number {
		fmt.Println(fmt.Sprintf("出牌 %v", card.translate()))
		p.uno.pool.Push(card)
		return true
	}
	return false
}
