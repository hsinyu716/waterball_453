package card

import (
	"fmt"
)

type CardShowdown struct {
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
	card := CardShowdown{}
	return &card
}

func (c *CardShowdown) GenerateDeck() []Card {
	suit := []SuitEnumType{SUITClub, SUITDiamond, SUITHeart, SUITSpade}
	rank := []RankEnumType{RANKTwo, RANKThree, RANKFore, RANKFive, RANKSix, RANKSeven, RANKEight, RANKNine, RANKTen, RANKJack, RANKQueen, RANKKing, RANKAce}
	var cards []Card
	for _, s := range suit {
		for _, r := range rank {
			cards = append(cards, &CardShowdown{r, s})
		}
	}
	return cards
}

func (c *CardShowdown) Translate() string {
	return fmt.Sprintf("%s %s", suits[c.GetSuit()], ranks[c.GetRank()])
}

func (c *CardShowdown) setRank(rank RankEnumType) {
	c.rank = rank
}

func (c *CardShowdown) setSuit(suit SuitEnumType) {
	c.suit = suit
}

func (c *CardShowdown) GetRank() RankEnumType {
	return c.rank
}

func (c *CardShowdown) GetSuit() SuitEnumType {
	return c.suit
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
