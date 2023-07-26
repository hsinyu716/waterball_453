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
}

type PlayerService interface {
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
	ShowCard() *Card
	GainPoint()
	GetPoint() int
	GetName() string
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

func (p *PlayerAdapter) ShowCard() *Card {
	return p.GetHand().Show()
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

func (p *PlayerAdapter) filterOtherPlayer() []PlayerService {
	var selectPlayers []PlayerService
	for _, player := range p.showdown.GetPlayers() {
		if p.name != player.GetName() {
			selectPlayers = append(selectPlayers, player)
		}
	}
	return selectPlayers
}

func (p *PlayerAdapter) selectExchangeHandsTarget(players []PlayerService) *ExchangeHands {
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

func printPlayerChoices(players []PlayerService) {
	str := "\n"
	for i, player := range players {
		str += fmt.Sprintf("(%d) %s \n", i, player.GetName())
	}
	fmt.Println(fmt.Sprintf("Select the target %s", str))
}
