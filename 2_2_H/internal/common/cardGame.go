package common

import (
	"fmt"
)

type GameType interface {
	*CardUno | *Showdown
}

type ICardGame[T any] interface {
	nameThemselves()
	Start()
	tableTopCard() Card
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

type CardGame[T any] struct {
	game  ICardGame[T]
	Trash *Desk
	Desk  *Desk
}

func NewCardGame[T any](game ICardGame[T]) *CardGame[T] {
	return &CardGame[T]{
		game: game,
	}
}

func (c *CardGame[T]) Start() {
	c.nameThemselves()
	c.GetDesk().Shuffle()
	c.drawHand()

	c.showTable()

	c.playRound()
	c.gameOver()
}

func (c *CardGame[T]) GetDesk() *Desk {
	return c.game.GetDesk()
}

func (c *CardGame[T]) GetTrash() *Desk {
	return c.game.GetTrash()
}

func (c *CardGame[T]) SetPlayers(p []IPlayer) {
	c.game.SetPlayers(p)
}

func (c *CardGame[T]) GetPlayers() []IPlayer {
	return c.game.GetPlayers()
}

func (c *CardGame[T]) tableTopCard() Card {
	return c.GetTrash().TopCard()
}

func (c *CardGame[T]) nameThemselves() {
	for i, p := range c.GetPlayers() {
		p.SetGame(c.game)
		p.NameHimself(i + 1)
		p.SetHand(NewHand(p.GetName()))
	}
}

func (c *CardGame[T]) drawHand() {
	c.game.drawHand()
}

func (c *CardGame[T]) showTable() {
	c.game.showTable()
}

func (c *CardGame[T]) playRound() {
	c.game.playRound()
}

func (c *CardGame[T]) takeTurn(player IPlayer) {
	c.game.takeTurn(player)
}

func (c *CardGame[T]) compareToWinner() IPlayer {
	return c.game.compareToWinner()
}

func (c *CardGame[T]) gameOver() {
	var winner IPlayer
	winner = c.compareToWinner()
	fmt.Println(fmt.Sprintf("The winner is %s.\n", winner.GetName()))
}
