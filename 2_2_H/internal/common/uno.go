package common

import (
	"fmt"
)

const HandCardNumber = 5

type Uno struct {
	desk      *Desk
	pool      *Desk
	players   []IPlayer
	turnMoves []*TurnMove
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
}

func (u *Uno) GetPlayers() []IPlayer {
	return u.players
}

func (u *Uno) Start() {
	u.nameThemselves()
	u.desk.Shuffle()
	u.drawHand()
	u.showTable()
	u.playRound()
	u.gameOver()
}

func (u *Uno) nameThemselves() {
	for i, p := range u.players {
		fmt.Println(p)
		p.SetGame(u)
		p.NameHimself(i + 1)
		p.SetHand(NewHand(p.GetName()))
	}
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

func (u *Uno) tableTopCard() *Card {
	return u.pool.TopCard()
}

func (u *Uno) playRound() {
	i := 0
	for {
		fmt.Println(fmt.Sprintf("ROUND %d", i+1))
		for _, player := range u.players {
			u.takeTurn(player)
			if player.GetCardSize() == 0 {
				return
			}
		}
		i++
	}
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

func (u *Uno) gameOver() {
	var winner IPlayer
	winner = u.compareToWinner()
	fmt.Println(fmt.Sprintf("The winner is %s.\n", winner.GetName()))
}
