package card

import (
	"fmt"
)

type Showdown struct {
	rank RankEnumType
	suit SuitEnumType
}

type ICardShowdown interface {
	setRank(rank RankEnumType)
	setSuit(suit SuitEnumType)
	GetRank() RankEnumType
	GetSuit() SuitEnumType
}

func NewCardShowdown() Card {
	card := Showdown{}
	return &card
}

func (s *Showdown) InitDeck() []Card {
	suit := []SuitEnumType{SUITClub, SUITDiamond, SUITHeart, SUITSpade}
	rank := []RankEnumType{RANKTwo, RANKThree, RANKFore, RANKFive, RANKSix, RANKSeven, RANKEight, RANKNine, RANKTen, RANKJack, RANKQueen, RANKKing, RANKAce}
	var cards []Card
	for _, su := range suit {
		for _, r := range rank {
			cards = append(cards, &Showdown{r, su})
		}
	}
	return cards
}

func (s *Showdown) Translate() string {
	return fmt.Sprintf("%s %s", suits[s.GetSuit()], ranks[s.GetRank()])
}

func (s *Showdown) setRank(rank RankEnumType) {
	s.rank = rank
}

func (s *Showdown) setSuit(suit SuitEnumType) {
	s.suit = suit
}

func (s *Showdown) GetRank() RankEnumType {
	return s.rank
}

func (s *Showdown) GetSuit() SuitEnumType {
	return s.suit
}

type RankEnumType int

type SuitEnumType int

const (
	RANKTwo RankEnumType = iota
	RANKThree
	RANKFore
	RANKFive
	RANKSix
	RANKSeven
	RANKEight
	RANKNine
	RANKTen
	RANKJack
	RANKQueen
	RANKKing
	RANKAce
)

const (
	SUITClub SuitEnumType = iota
	SUITDiamond
	SUITHeart
	SUITSpade
)

var suits = map[SuitEnumType]string{
	SUITClub:    "梅花",
	SUITDiamond: "方塊",
	SUITHeart:   "紅心",
	SUITSpade:   "黑桃",
}

var ranks = map[RankEnumType]string{
	RANKTwo:   "2",
	RANKThree: "3",
	RANKFore:  "4",
	RANKFive:  "5",
	RANKSix:   "6",
	RANKSeven: "7",
	RANKEight: "8",
	RANKNine:  "9",
	RANKTen:   "10",
	RANKJack:  "J",
	RANKQueen: "Q",
	RANKKing:  "K",
	RANKAce:   "A",
}
