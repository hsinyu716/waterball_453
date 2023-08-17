package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type PlayerAdapter[T any] struct {
	name          string
	point         int
	showdown      *Showdown
	hand          *Hand
	exchangeHands *ExchangeHands
	game          ICardGame[T]
}

type IPlayer interface {
	NameHimself(i int)
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

	SetGame(game ICardGame[any])
	GetGame() ICardGame[any]
	TakeTurnUno()
}

func (p *PlayerAdapter[T]) NameHimself(i int) {
}

func (p *PlayerAdapter[T]) SetHand(hand *Hand) {
	p.hand = hand
}

func (p *PlayerAdapter[T]) AddHandCard(card *Card) {
	p.hand.AddCard(card)
}

func (p *PlayerAdapter[T]) GetHand() *Hand {
	return p.hand
}

func (p *PlayerAdapter[T]) TakeTurn() *TurnMove {
	return nil
}

func (p *PlayerAdapter[T]) MakeExchangeHandsDecision() *ExchangeHands {
	return nil
}

func (p *PlayerAdapter[T]) HasUsedExchangeHands() bool {
	return p.exchangeHands != nil
}

func (p *PlayerAdapter[T]) GetExchangeHands() *ExchangeHands {
	return p.exchangeHands
}

func (p *PlayerAdapter[T]) SetExchangeHands(exchangeHands *ExchangeHands) {
	p.exchangeHands = exchangeHands
}

func (p *PlayerAdapter[T]) ShowCard(index int) *Card {
	return p.GetHand().Show(index)
}

func (p *PlayerAdapter[T]) GainPoint() {
	p.point++
}

func (p *PlayerAdapter[T]) GetPoint() int {
	return p.point
}

func (p *PlayerAdapter[T]) GetName() string {
	return p.name
}

func (p *PlayerAdapter[T]) GetCardSize() int {
	return len(p.GetHand().cards)
}

func (p *PlayerAdapter[T]) filterOtherPlayer() []IPlayer {
	var selectPlayers []IPlayer
	for _, player := range p.game.GetPlayers() {
		if p.name != player.GetName() {
			selectPlayers = append(selectPlayers, player)
		}
	}
	return selectPlayers
}

func (p *PlayerAdapter[T]) selectExchangeHandsTarget(players []IPlayer) *ExchangeHands {
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

func (p *PlayerAdapter[T]) SetGame(game ICardGame[any]) {
	p.game = game
}

func (p *PlayerAdapter[T]) GetGame() ICardGame[any] {
	return p.game
}

func (p *PlayerAdapter[T]) TakeTurnUno() {
	topCard := p.game.tableTopCard()
	fmt.Println(fmt.Sprintf("topCard %v", topCard.translate()))
	for i, card := range p.GetHand().cards {
		if p.compareCard(topCard, card) {
			p.ShowCard(i)
			return
		}
	}
	if p.game.GetDesk().Size() == 0 {
		for _, c := range p.game.GetPool().cards {
			p.game.GetDesk().Push(c)
		}
		p.game.GetDesk().Shuffle()
		p.game.GetPool().cards = nil
		p.game.GetPool().Push(topCard)
	}
	card := p.game.GetDesk().DrawCard()
	fmt.Println(fmt.Sprintf("抽卡 %v", card.translate()))
	// 抽卡判斷可以出
	if p.compareCard(topCard, card) {
		return
	}
	p.AddHandCard(card)
	return
}

func (p *PlayerAdapter[T]) compareCard(topCard *Card, card *Card) bool {
	if topCard.color == card.color || topCard.number == card.number {
		fmt.Println(fmt.Sprintf("出牌 %v", card.translate()))
		p.game.GetPool().Push(card)
		return true
	}
	return false
}
