package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

const RoundNum = 13

type Showdown struct {
	CardGame
	turnMoves []*TurnMove
}

func NewShowdown(players *[]IPlayer) *Showdown {
	return &Showdown{
		CardGame: CardGame{
			deck:    NewDeck(card.NewCardShowdown().InitDeck()),
			players: *players,
		},
	}
}

func (s *Showdown) GetPlayers() []IPlayer {
	return s.players
}

func (s *Showdown) GetDeck() *Deck {
	return s.deck
}

func (s *Showdown) getHandLimit() int {
	return RoundNum
}

func (s *Showdown) takeTurn(player IPlayer) {
	fmt.Println(fmt.Sprintf("It's (%s)'s turn", player.GetName()))
	turnMove := player.TakeTurn()
	s.turnMoves = append(s.turnMoves, turnMove)
}

func (s *Showdown) checkWinner(winner, player IPlayer) bool {
	return winner.GetPoint() > player.GetPoint()
}

func (s *Showdown) checkOver(i int, _ IPlayer) (bool, bool) {
	return i == RoundNum-1, false
}

func (s *Showdown) winnerInRound() {
	s.printShowCards()
	winnerTurnMove := s.compareToTurn()
	winner := winnerTurnMove.GetPlayer()
	winner.GainPoint()
	fmt.Println(fmt.Sprintf("%s wins this round.\n", winner.GetName()))
}

func (s *Showdown) printShowCards() {
	str := "Show cards: "
	for _, move := range s.turnMoves {
		card0 := move.GetPlayer().GetHand().Show(0)
		move.SetShowCard(card0)
		str += fmt.Sprintf("%v ", move.GetShowCard().Translate())
	}
	fmt.Println(str)
}

func (s *Showdown) compareToTurn() *TurnMove {
	winnerTurnMove := s.turnMoves[0]
	for _, move := range s.turnMoves {
		moveCard := move.GetShowCard().(*card.Showdown)
		winnerCard := winnerTurnMove.GetShowCard().(*card.Showdown)
		if moveCard.GetRank() > winnerCard.GetRank() {
			winnerTurnMove = move
		} else if winnerCard.GetRank() == moveCard.GetRank() {
			if moveCard.GetSuit() > winnerCard.GetSuit() {
				winnerTurnMove = move
			}
		}
	}
	return winnerTurnMove
}

func (s *Showdown) cleanTurn() {
	s.turnMoves = nil
}

func (s *Showdown) showTable() {
	// do nothing
}
