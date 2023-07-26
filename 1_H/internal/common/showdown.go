package common

import (
	"fmt"
)

const RoundNum = 13

type Showdown struct {
	desk      *Desk
	players   []PlayerService
	turnMoves []*TurnMove
}

func NewShowdown(desk *Desk, players *[]PlayerService) *Showdown {
	return &Showdown{
		desk:    desk,
		players: *players,
	}
}

type ShowdownService interface {
	GetPlayers() []PlayerService
	Start()
}

func (s *Showdown) GetPlayers() []PlayerService {
	return s.players
}

func (s *Showdown) Start() {
	s.nameThemselves()
	s.desk.Shuffle()
	s.drawHand()
	s.playRound()
	s.gameOver()
}

func (s *Showdown) nameThemselves() {
	for i, p := range s.players {
		p.SetShowdown(s)
		p.NameHimself(i + 1)
		p.SetHand(NewHand(p.GetName()))
	}
}

func (s *Showdown) drawHand() {
	size := s.desk.Size()
	for i := 0; i < size; i++ {
		card := s.desk.DrawCard()
		s.players[i%4].AddHandCard(card)
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

func (s *Showdown) takeTurn(player PlayerService) {
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
		card := move.GetPlayer().GetHand().Show()
		move.SetShowCard(card)
		str += fmt.Sprintf("%v", move.GetShowCard())
	}
	fmt.Println(str)
}

func (s *Showdown) compareToTurn() *TurnMove {
	winnerTurnMove := s.turnMoves[0]
	for _, move := range s.turnMoves {
		if winnerTurnMove.GetShowCard().GetRank() > move.GetShowCard().GetRank() {
			winnerTurnMove = move
		} else if winnerTurnMove.GetShowCard().GetRank() == move.GetShowCard().GetRank() {
			if winnerTurnMove.GetShowCard().GetSuit() > move.GetShowCard().GetSuit() {
				winnerTurnMove = move
			}
		}
	}
	return winnerTurnMove
}

func (s *Showdown) compareToWinner() PlayerService {
	players := s.GetPlayers()
	winner := players[0]
	for _, player := range players {
		if winner.GetPoint() > player.GetPoint() {
			winner = player
		}
	}
	return winner
}

func (s *Showdown) gameOver() {
	var winner PlayerService
	winner = s.compareToWinner()
	fmt.Println(fmt.Sprintf("The winner is %s.\n", winner.GetName()))
}
