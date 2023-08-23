package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

type ICardGame interface {
	NameThemselves()
	Start()
	tableTopCard() card.Card
	GetDesk() *Desk
	GetTrash() *Desk
	SetPlayers(p []IPlayer)
	GetPlayers() []IPlayer
	drawHand()
	showTable()
	takeTurn(player IPlayer)
	playRound()
	compareToWinner() IPlayer
}

type CardGame struct {
	game  ICardGame
	trash *Desk
	desk  *Desk
}

func NewCardGame(game ICardGame) *CardGame {
	return &CardGame{
		game: game,
	}
}

func (c *CardGame) Start() {
	c.NameThemselves()
	c.GetDesk().Shuffle()
	c.drawHand()

	c.showTable()

	c.playRound()
	c.gameOver()
}

func (c *CardGame) GetDesk() *Desk {
	return c.game.GetDesk()
}

func (c *CardGame) GetTrash() *Desk {
	return c.game.GetTrash()
}

func (c *CardGame) SetPlayers(p []IPlayer) {
	c.game.SetPlayers(p)
}

func (c *CardGame) GetPlayers() []IPlayer {
	return c.game.GetPlayers()
}

func (c *CardGame) tableTopCard() card.Card {
	return c.GetTrash().TopCard()
}

func (c *CardGame) NameThemselves() {
	for i, p := range c.GetPlayers() {
		p.SetGame(c.game)
		p.NameHimself(i + 1)
		p.SetHand(NewHand(p.GetName()))
	}
}

func (c *CardGame) drawHand() {
	c.game.drawHand()
}

func (c *CardGame) showTable() {
	c.game.showTable()
}

func (c *CardGame) playRound() {
	c.game.playRound()
}

func (c *CardGame) takeTurn(player IPlayer) {
	c.game.takeTurn(player)
}

func (c *CardGame) compareToWinner() IPlayer {
	return c.game.compareToWinner()
}

func (c *CardGame) gameOver() {
	var winner IPlayer
	winner = c.compareToWinner()
	fmt.Println(fmt.Sprintf("The winner is %s.\n", winner.GetName()))
}
