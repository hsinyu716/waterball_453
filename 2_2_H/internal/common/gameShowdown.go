package common

import (
	"fmt"
)

const RoundNum = 13

type GameShowdown struct {
	CardGame[*GameShowdown]
	desk      *Desk
	players   []IPlayer
	turnMoves []*TurnMove
}

func NewShowdown(players *[]IPlayer) *GameShowdown {
	return &GameShowdown{
		desk:    NewDesk(NewCardShowdown().GenerateDeck()),
		players: *players,
	}
}

func (s *GameShowdown) GetPlayers() []IPlayer {
	return s.players
}

func (s *GameShowdown) GetDesk() *Desk {
	return s.desk
}

func (s *GameShowdown) drawHand() {
	size := s.desk.Size()
	for i := 0; i < size; i++ {
		card := s.desk.DrawCard()
		if s.players[i%4].GetCardSize() > 13 {
			panic("over 13")
		}
		s.players[i%4].AddHandCard(card)
	}
}

func (s *GameShowdown) playRound() {
	for i := 0; i < RoundNum; i++ {
		fmt.Println(fmt.Sprintf("ROUND %d", i+1))
		for _, player := range s.players {
			s.takeTurn(player)
		}
		s.showdown()
		s.turnMoves = nil
	}
}

func (s *GameShowdown) takeTurn(player IPlayer) {
	fmt.Println(fmt.Sprintf("It's (%s)'s turn", player.GetName()))
	turnMove := player.TakeTurn()
	s.turnMoves = append(s.turnMoves, turnMove)
}

func (s *GameShowdown) showdown() {
	s.printShowCards()
	winnerTurnMove := s.compareToTurn()
	winner := winnerTurnMove.GetPlayer()
	winner.GainPoint()
	fmt.Println(fmt.Sprintf("%s wins this round.\n", winner.GetName()))
}

func (s *GameShowdown) printShowCards() {
	str := "Show cards: "
	for _, move := range s.turnMoves {
		card := move.GetPlayer().GetHand().Show(0)
		move.SetShowCard(card)
		str += fmt.Sprintf("%v ", move.GetShowCard().Translate())
	}
	fmt.Println(str)
}

func (s *GameShowdown) compareToTurn() *TurnMove {
	winnerTurnMove := s.turnMoves[0]
	//for _, move := range s.turnMoves {
	//	if move.GetShowCard().GetRank() > winnerTurnMove.GetShowCard().GetRank() {
	//		winnerTurnMove = move
	//	} else if winnerTurnMove.GetShowCard().GetRank() == move.GetShowCard().GetRank() {
	//		if move.GetShowCard().GetSuit() > winnerTurnMove.GetShowCard().GetSuit() {
	//			winnerTurnMove = move
	//		}
	//	}
	//}
	return winnerTurnMove
}

func (s *GameShowdown) compareToWinner() IPlayer {
	players := s.GetPlayers()
	winner := players[0]
	for _, player := range players {
		if winner.GetPoint() > player.GetPoint() {
			winner = player
		}
	}
	return winner
}

func (s *GameShowdown) showTable() {
}
