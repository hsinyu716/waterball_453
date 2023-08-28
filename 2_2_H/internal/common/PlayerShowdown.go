package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

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

func (p *PlayerAdapter) TakeTurn() *TurnMove {
	turnMove := NewTurnMove(p, nil, nil)
	return turnMove
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

func (p *PlayerAdapter) GainPoint() {
	p.point++
}

func (p *PlayerAdapter) GetPoint() int {
	return p.point
}

func (p *PlayerAdapter) FilterOtherPlayer() []IPlayer {
	var selectPlayers []IPlayer
	for _, player := range p.game.GetPlayers() {
		if p.name != player.GetName() {
			selectPlayers = append(selectPlayers, player)
		}
	}
	return selectPlayers
}

func (p *PlayerAdapter) SelectExchangeHandsTarget(players []IPlayer) *ExchangeHands {
	printPlayerChoices(players)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	input = strings.TrimSuffix(input, "\n")
	targetIndex, err := strconv.Atoi(input)
	if targetIndex >= len(players) || targetIndex < 0 {
		return p.SelectExchangeHandsTarget(players)
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
