package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

const RoundNum = 13

type Showdown struct {
	CardGame
	desk      *Desk
	players   []IPlayer
	turnMoves []*TurnMove
}

func NewShowdown(players *[]IPlayer) *Showdown {
	return &Showdown{
		desk:    NewDesk(card.NewCardShowdown().GenerateDeck()),
		players: *players,
	}
}

func (s *Showdown) GetPlayers() []IPlayer {
	return s.players
}

func (s *Showdown) GetDesk() *Desk {
	return s.desk
}

func (s *Showdown) drawHand() {
	size := s.desk.Size()
	for i := 0; i < size; i++ {
		card0 := s.desk.DrawCard()
		if s.players[i%4].GetCardSize() > 13 {
			panic("over 13")
		}
		s.players[i%4].AddHandCard(card0)
	}
}

func (s *Showdown) playRound() {
	for i := 0; i < RoundNum; i++ {
		fmt.Println(fmt.Sprintf("ROUND %d", i+1))
		for _, player := range s.players {
			s.takeTurn(player)
		}
		s.showdown()
		s.turnMoves = nil
	}
}

func (s *Showdown) takeTurn(player IPlayer) {
	fmt.Println(fmt.Sprintf("It's (%s)'s turn", player.GetName()))
	turnMove := player.TakeTurn()
	s.turnMoves = append(s.turnMoves, turnMove)
}

func (s *Showdown) showdown() {
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
		moveCard := move.GetShowCard().(*card.CardShowdown)
		winnerCard := winnerTurnMove.GetShowCard().(*card.CardShowdown)
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

func (s *Showdown) compareToWinner() IPlayer {
	players := s.GetPlayers()
	winner := players[0]
	for _, player := range players {
		if winner.GetPoint() > player.GetPoint() {
			winner = player
		}
	}
	return winner
}

func (s *Showdown) showTable() {
	// do nothing
}
