package common

import (
	"fmt"
)

const HandCardNumber = 5

type GameUno struct {
	CardGame[*GameUno]
	players []IPlayer
}

func NewGameUno(players *[]IPlayer) *GameUno {
	return &GameUno{
		CardGame: CardGame[*GameUno]{
			Desk:  NewDesk(NewCardUno().GenerateDeck()),
			Trash: NewDesk([]Card{}),
		},
		players: *players,
	}
}

//func (u *CardUno) Start() {
//	u.nameThemselves()
//	u.GetDesk().Shuffle()
//	u.drawHand()
//
//	u.showTable()
//
//	u.playRound()
//	u.gameOver()
//}

func (g *GameUno) GetPlayers() []IPlayer {
	return g.players
}

func (g *GameUno) GetDesk() *Desk {
	return g.Desk
}

func (g *GameUno) GetTrash() *Desk {
	return g.Trash
}

func (g *GameUno) drawHand() {
	size := g.Desk.Size()
	for i := 0; i < size; i++ {
		if i == len(g.players)*HandCardNumber {
			break
		}
		card := g.Desk.DrawCard()
		g.players[i%4].AddHandCard(card)
	}
}

func (g *GameUno) showTable() {
	card := g.Desk.DrawCard()
	g.Trash.Push(card)
	fmt.Println(fmt.Sprintf("First card is %v", card.Translate()))
}

func (g *GameUno) playRound() {
	i := 0
	for {
		fmt.Println(fmt.Sprintf("ROUND %d", i+1))
		for _, player := range g.GetPlayers() {
			g.takeTurn(player)
			if player.GetCardSize() == 0 {
				return
			}
		}
		i++
	}
}

func (g *GameUno) tableTopCard() Card {
	return g.Trash.TopCard()
}

func (g *GameUno) takeTurn(player IPlayer) {
	fmt.Println(fmt.Sprintf("It's (%s)'s turn", player.GetName()))
	player.TakeTurnUno()
	if player.GetCardSize() == 1 {
		fmt.Println(fmt.Sprintf("=======%s UNO~", player.GetName()))
	}
}

func (g *GameUno) compareToWinner() IPlayer {
	players := g.GetPlayers()
	winner := players[0]
	for _, player := range players {
		if player.GetCardSize() == 0 {
			winner = player
		}
	}
	return winner
}
