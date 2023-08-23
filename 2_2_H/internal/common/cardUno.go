package common

import (
	"fmt"
)

type CardUno struct {
	number NumberEnumType
	color  ColorEnumType
}

type ICardUno interface {
	setColor(color ColorEnumType)
	setNumber(number NumberEnumType)
	getColor() ColorEnumType
	getNumber() NumberEnumType
	CompareCard(topCard CardUno) bool
}

func NewCardUno() Card {
	card := CardUno{}
	return &card
}

func (c *CardUno) Translate() string {
	return fmt.Sprintf("%s %d", colors[c.getColor()], c.getNumber()+1)
}

func (c *CardUno) GenerateDeck() []Card {
	color := []ColorEnumType{COLORBlue, COLORRed, COLORYellow, COLORGreen}
	number := []NumberEnumType{N0, N1, N2, N3, N4, N5, N6, N7, N8, N9}
	var cards []Card
	for _, s := range color {
		for _, r := range number {
			cards = append(cards, &CardUno{r, s})
		}
	}
	return cards
}

func (c *CardUno) CompareCard(topCard Card) bool {
	c0 := topCard.(*CardUno)
	if c0.color == c.color || c0.number == c.number {
		fmt.Println(fmt.Sprintf("出牌 %v", c.Translate()))
		return true
	}
	return false
}

func (c *CardUno) setColor(color ColorEnumType) {
	c.color = color
}

func (c *CardUno) setNumber(number NumberEnumType) {
	c.number = number
}

func (c *CardUno) getColor() ColorEnumType {
	return c.color
}

func (c *CardUno) getNumber() NumberEnumType {
	return c.number
}

type ColorEnumType int

type NumberEnumType int

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
	COLORBlue ColorEnumType = iota
	COLORRed
	COLORYellow
	COLORGreen
)

var colors = map[ColorEnumType]string{
	COLORBlue:   "藍色",
	COLORRed:    "紅色",
	COLORYellow: "黃色",
	COLORGreen:  "綠色",
}
