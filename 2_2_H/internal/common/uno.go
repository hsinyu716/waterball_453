package common

import (
	"fmt"
)

const HandCardNumber = 5

type Uno struct {
	CardsGame[*Uno]
	desk    *Desk
	pool    *Desk
	players []IPlayer
}

func NewUno(desk *Desk, players *[]IPlayer) *Uno {
	return &Uno{
		desk:    desk,
		pool:    NewDesk(),
		players: *players,
	}
}

type IUno interface {
	GetPlayers() []IPlayer
	Start()
	GetDesk() *Desk
}

func (u *Uno) GetPlayers() []IPlayer {
	return u.players
}

func (u *Uno) GetDesk() *Desk {
	return u.desk
}

func (u *Uno) GetPool() *Desk {
	return u.pool
}

func (u *Uno) drawHand() {
	size := u.desk.Size()
	for i := 0; i < size; i++ {
		if i == len(u.players)*HandCardNumber {
			break
		}
		card := u.desk.DrawCard()
		u.players[i%4].AddHandCard(card)
	}
}

func (u *Uno) showTable() {
	card := u.desk.DrawCard()
	u.pool.Push(card)
	fmt.Println(fmt.Sprintf("First card is %v", card.translate()))
}

func (u *Uno) playRound() {
	i := 0
	for {
		fmt.Println(fmt.Sprintf("ROUND %d", i+1))
		for _, player := range u.GetPlayers() {
			u.takeTurn(player)
			if player.GetCardSize() == 0 {
				return
			}
		}
		i++
	}
}

func (u *Uno) tableTopCard() *Card {
	return u.pool.TopCard()
}

func (u *Uno) takeTurn(player IPlayer) {
	fmt.Println(fmt.Sprintf("It's (%s)'s turn", player.GetName()))
	player.TakeTurnUno()
	if player.GetCardSize() == 1 {
		fmt.Println(fmt.Sprintf("=======%s UNO~", player.GetName()))
	}
}

func (u *Uno) compareToWinner() IPlayer {
	players := u.GetPlayers()
	winner := players[0]
	for _, player := range players {
		if player.GetCardSize() == 0 {
			winner = player
		}
	}
	return winner
}
