package common

import (
	"fmt"
)

type GameType interface {
	*Uno | *Showdown
}

type ICardGame[T any] interface {
	nameThemselves()
	Start()
	tableTopCard() *Card
	GetDesk() *Desk
	GetPool() *Desk
	SetPlayers(p []IPlayer)
	GetPlayers() []IPlayer
	drawHand()
	showTable()
	takeTurn(player IPlayer)
	playRound()
	compareToWinner() IPlayer
}

type CardsGame[T any] struct {
	game ICardGame[T]
	pool *Desk
	desk *Desk
}

func NewCardsGame[T any](game ICardGame[T]) *CardsGame[T] {
	return &CardsGame[T]{
		game: game,
	}
}

func (c *CardsGame[T]) GetDesk() *Desk {
	return c.game.GetDesk()
}
func (c *CardsGame[T]) GetPool() *Desk {
	return c.game.GetPool()
}

func (c *CardsGame[T]) SetPlayers(p []IPlayer) {
	c.game.SetPlayers(p)
}

func (c *CardsGame[T]) GetPlayers() []IPlayer {
	return c.game.GetPlayers()
}

func (c *CardsGame[T]) tableTopCard() *Card {
	return c.GetPool().TopCard()
}

func (c *CardsGame[T]) nameThemselves() {
	for i, p := range c.GetPlayers() {
		p.SetGame(c.game)
		p.NameHimself(i + 1)
		p.SetHand(NewHand(p.GetName()))
	}
}

func (c *CardsGame[T]) Start() {
	c.nameThemselves()
	c.GetDesk().Shuffle()
	c.drawHand()

	c.showTable()

	c.playRound()
	c.gameOver()
}

func (c *CardsGame[T]) drawHand() {
	c.game.drawHand()
}

func (c *CardsGame[T]) showTable() {
}

func (c *CardsGame[T]) playRound() {
	c.game.playRound()
}

func (c *CardsGame[T]) takeTurn(player IPlayer) {
	c.game.takeTurn(player)
}

func (c *CardsGame[T]) compareToWinner() IPlayer {
	return c.game.compareToWinner()
}

func (c *CardsGame[T]) gameOver() {
	var winner IPlayer
	winner = c.compareToWinner()
	fmt.Println(fmt.Sprintf("The winner is %s.\n", winner.GetName()))
}
