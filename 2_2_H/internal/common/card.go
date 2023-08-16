package common

import "fmt"

type Card struct {
	rank   RankEnumType
	suit   SuitEnumType
	number NumberEnumType
	color  ColorEnumType
}

func NewCard(rank RankEnumType, suit SuitEnumType) *Card {
	return &Card{
		rank: rank,
		suit: suit,
	}
}

func NewCard2(number NumberEnumType, color ColorEnumType) *Card {
	return &Card{
		number: number,
		color:  color,
	}
}

func (c *Card) GetRank() RankEnumType {
	return c.rank
}

func (c *Card) GetSuit() SuitEnumType {
	return c.suit
}

func (c *Card) GetColor() ColorEnumType {
	return c.color
}

func (c *Card) GetNumber() NumberEnumType {
	return c.number
}

type RankEnumType int

type SuitEnumType int

type ColorEnumType int

type NumberEnumType int

const (
	TWO RankEnumType = iota
	THREE
	FORE
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

const (
	Club SuitEnumType = iota
	Diamond
	Heart
	Spade
)

const (
	N0 NumberEnumType = iota
	N1
	N2
	N3
	N4
	N5
	N6
	N7
	N8
	N9
)

const (
	BLUE ColorEnumType = iota
	RED
	YELLOW
	GREEN
)

var colors = map[ColorEnumType]string{
	BLUE:   "藍色",
	RED:    "紅色",
	YELLOW: "黃色",
	GREEN:  "綠色",
}

func (c *Card) translate() string {
	return fmt.Sprintf("%s %d", colors[c.color], c.GetNumber()+1)
}
