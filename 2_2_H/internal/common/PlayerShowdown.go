package common

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayerShowdown struct {
	point         int
	exchangeHands *ExchangeHands
}

type IPlayerShowdown interface {
	TakeTurn() *TurnMove
	MakeExchangeHandsDecision() *ExchangeHands
	HasUsedExchangeHands() bool
	GetExchangeHands() *ExchangeHands
	SetExchangeHands(exchangeHands *ExchangeHands)
	GainPoint()
	GetPoint() int
}

func (p *PlayerAdapter[T]) TakeTurn() *TurnMove {
	turnMove := NewTurnMove(p, nil, nil)
	return turnMove
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

func (p *PlayerAdapter[T]) GainPoint() {
	p.point++
}

func (p *PlayerAdapter[T]) GetPoint() int {
	return p.point
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
