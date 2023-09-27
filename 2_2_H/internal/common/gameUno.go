package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

const HandCardNumber = 5

type Uno struct {
	CardGame
}

func NewGameUno(players *[]IPlayer) *Uno {
	return &Uno{
		CardGame: CardGame{
			deck:    NewDeck(card.NewCardUno().InitDeck()),
			trash:   NewDeck([]card.Card{}),
			players: *players,
		},
	}
}

func (u *Uno) GetPlayers() []IPlayer {
	return u.players
}

func (u *Uno) GetDeck() *Deck {
	return u.deck
}

func (u *Uno) GetTrash() *Deck {
	return u.trash
}

func (u *Uno) getHandLimit() int {
	return HandCardNumber
}

func (u *Uno) showTable() {
	card0 := u.deck.DrawCard()
	u.trash.Push(card0)
	fmt.Println(fmt.Sprintf("First card is %v", card0.Translate()))
}

func (u *Uno) checkOver(_ int, player IPlayer) (bool, bool) {
	return player.GetCardSize() == 0, true
}

func (u *Uno) tableTopCard() card.Card {
	return u.trash.TopCard()
}

func (u *Uno) takeTurn(player IPlayer) {
	fmt.Println(fmt.Sprintf("It's (%s)'s turn", player.GetName()))
	player.TakeTurnUno()
	if player.GetCardSize() == 1 {
		fmt.Println(fmt.Sprintf("=======%s UNO~~", player.GetName()))
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

func (u *Uno) checkWinner(_, player IPlayer) bool {
	return player.GetCardSize() == 0
}

func (u *Uno) cleanTurn() {
	// do nothing
}

func (u *Uno) winnerInRound() {
	// do nothing
}
