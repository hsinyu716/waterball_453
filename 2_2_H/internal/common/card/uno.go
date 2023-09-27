package card

import (
	"fmt"
)

type Uno struct {
	number NumberEnumType
	color  ColorEnumType
}

func NewCardUno() Card {
	card := Uno{}
	return &card
}

func (u *Uno) InitDeck() []Card {
	color := []ColorEnumType{COLORBlue, COLORRed, COLORYellow, COLORGreen}
	number := []NumberEnumType{N0, N1, N2, N3, N4, N5, N6, N7, N8, N9}
	var cards []Card
	for _, s := range color {
		for _, r := range number {
			cards = append(cards, &Uno{r, s})
		}
	}
	return cards
}

func (u *Uno) Translate() string {
	return fmt.Sprintf("%s %d", colors[u.getColor()], u.getNumber()+1)
}

func (u *Uno) CompareCard(topCard Card) bool {
	c0 := topCard.(*Uno)
	if c0.color == u.color || c0.number == u.number {
		fmt.Println(fmt.Sprintf("出牌 %v", u.Translate()))
		return true
	}
	return false
}

func (u *Uno) setColor(color ColorEnumType) {
	u.color = color
}

func (u *Uno) setNumber(number NumberEnumType) {
	u.number = number
}

func (u *Uno) getColor() ColorEnumType {
	return u.color
}

func (u *Uno) getNumber() NumberEnumType {
	return u.number
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
