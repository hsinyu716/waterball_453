package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

type ICardGameStrategy interface {
	NameThemselves()
	tableTopCard() card.Card
	GetDeck() *Deck
	GetTrash() *Deck
	SetPlayers(p []IPlayer)
	GetPlayers() []IPlayer
	drawHand()
	showTable()
	takeTurn(player IPlayer)
	playRound()
	winnerInRound()
	checkOver(i int, player IPlayer) (bool, bool)
	compareToWinner() IPlayer
	checkWinner(winner, player IPlayer) bool
	cleanTurn()
	getHandLimit() int
}

type CardGame struct {
	gameStrategy ICardGameStrategy
	trash        *Deck
	deck         *Deck
	players      []IPlayer
}

func NewCardGame(game ICardGameStrategy) *CardGame {
	return &CardGame{
		gameStrategy: game,
	}
}

func (c *CardGame) Start() {
	c.NameThemselves()
	c.gameStrategy.GetDeck().Shuffle()
	c.drawHand()

	c.gameStrategy.showTable()

	c.playRound()
	c.gameOver()
}

func (c *CardGame) GetTrash() *Deck {
	return c.gameStrategy.GetTrash()
}

func (c *CardGame) SetPlayers(p []IPlayer) {
	c.gameStrategy.SetPlayers(p)
}

func (c *CardGame) GetPlayers() []IPlayer {
	return c.gameStrategy.GetPlayers()
}

func (c *CardGame) tableTopCard() card.Card {
	return c.GetTrash().TopCard()
}

func (c *CardGame) NameThemselves() {
	for i, p := range c.GetPlayers() {
		p.SetGame(c.gameStrategy)
		p.SetName(p.NameHimself(i + 1))
		p.SetHand(NewHand(p.GetName()))
	}
}

func (c *CardGame) drawHand() {
	size := c.gameStrategy.GetDeck().Size()
	for i := 0; i < size; i++ {
		card0 := c.gameStrategy.GetDeck().DrawCard()
		if c.gameStrategy.GetPlayers()[i%4].GetCardSize() > c.gameStrategy.getHandLimit() {
			break
		}
		c.gameStrategy.GetPlayers()[i%4].AddHandCard(card0)
	}
}

func (c *CardGame) playRound() {
	i := 0
	end := false
	breakLoop := false
	for !end {
		fmt.Println(fmt.Sprintf("ROUND %d", i+1))
		for _, player := range c.GetPlayers() {
			c.takeTurn(player)
			end, breakLoop = c.gameStrategy.checkOver(i, player)
			if end && breakLoop {
				break
			}
		}
		c.gameStrategy.winnerInRound()
		c.gameStrategy.cleanTurn()
		i++
	}
}

func (c *CardGame) takeTurn(player IPlayer) {
	c.gameStrategy.takeTurn(player)
}

func (c *CardGame) gameOver() {
	var winner IPlayer
	winner = c.compareToWinner()
	fmt.Println(fmt.Sprintf("The winner is %s.\n", winner.GetName()))
}

func (c *CardGame) compareToWinner() IPlayer {
	players := c.GetPlayers()
	winner := players[0]
	for _, player := range players {
		if c.gameStrategy.checkWinner(winner, player) {
			winner = player
		}
	}
	return winner
}
