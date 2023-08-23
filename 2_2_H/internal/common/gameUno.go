package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

const HandCardNumber = 5

type Uno struct {
	CardGame
	players []IPlayer
}

func NewGameUno(players *[]IPlayer) *Uno {
	return &Uno{
		CardGame: CardGame{
			desk:  NewDesk(card.NewCardUno().GenerateDeck()),
			trash: NewDesk([]card.Card{}),
		},
		players: *players,
	}
}

func (g *Uno) GetPlayers() []IPlayer {
	return g.players
}

func (g *Uno) GetDesk() *Desk {
	return g.desk
}

func (g *Uno) GetTrash() *Desk {
	return g.trash
}

func (g *Uno) drawHand() {
	size := g.desk.Size()
	for i := 0; i < size; i++ {
		if i == len(g.players)*HandCardNumber {
			break
		}
		card0 := g.desk.DrawCard()
		g.players[i%4].AddHandCard(card0)
	}
}

func (g *Uno) showTable() {
	card0 := g.desk.DrawCard()
	g.trash.Push(card0)
	fmt.Println(fmt.Sprintf("First card is %v", card0.Translate()))
}

func (g *Uno) playRound() {
	i := 0
	end := false
	for !end {
		fmt.Println(fmt.Sprintf("ROUND %d", i+1))
		for _, player := range g.GetPlayers() {
			g.takeTurn(player)
			end = player.GetCardSize() == 0
		}
		i++
	}
}

func (g *Uno) tableTopCard() card.Card {
	return g.trash.TopCard()
}

func (g *Uno) takeTurn(player IPlayer) {
	fmt.Println(fmt.Sprintf("It's (%s)'s turn", player.GetName()))
	player.TakeTurnUno()
	if player.GetCardSize() == 1 {
		fmt.Println(fmt.Sprintf("=======%s UNO~", player.GetName()))
	}
}

func (g *Uno) compareToWinner() IPlayer {
	players := g.GetPlayers()
	winner := players[0]
	for _, player := range players {
		if player.GetCardSize() == 0 {
			winner = player
		}
	}
	return winner
}
